#define PY_SSIZE_T_CLEAN
#include <Python.h>
#include <stdio.h>

/**
 * print_python_bytes - Prints information about Python bytes objects
 * @p: Pointer to the Python object
 */
void print_python_bytes(PyObject *p)
{
    PyBytesObject *bytes;
    Py_ssize_t size, i;
    char *string;

    printf("[.] bytes object info\n");

    if (!PyBytes_Check(p))
    {
        printf("  [ERROR] Invalid Bytes Object\n");
        return;
    }

    bytes = (PyBytesObject *)p;
    size = PyBytes_Size(p);
    string = PyBytes_AsString(p);

    printf("  size: %ld\n", size);
    printf("  trying string: %s\n", string);

    printf("  first 10 bytes:");
    for (i = 0; i < size && i < 10; ++i)
    {
        printf(" %02x", string[i] & 0xFF);
    }
    printf("\n");
}

/**
 * print_python_list - Prints information about Python lists
 * @p: Pointer to the Python object
 */
void print_python_list(PyObject *p)
{
    PyListObject *list;
    Py_ssize_t size, i;
    PyObject *element;

    printf("[*] Python list info\n");

    if (!PyList_Check(p))
    {
        printf("  [ERROR] Invalid List Object\n");
        return;
    }

    list = (PyListObject *)p;
    size = PyList_Size(p);

    printf("[*] Size of the Python List = %ld\n", size);
    printf("[*] Allocated = %ld\n", list->allocated);

    for (i = 0; i < size; ++i)
    {
        element = PyList_GetItem(p, i);
        printf("Element %ld: %s\n", i, Py_TYPE(element)->tp_name);

        if (PyBytes_Check(element))
        {
            print_python_bytes(element);
        }
    }
}

int main(void)
{
    PyObject *p;

    /* Test cases */
    p = PyBytes_FromString("Hello");
    print_python_bytes(p);
    Py_XDECREF(p);

    p = PyBytes_FromString("\xff\xf8\x00\x00\x00\x00\x00\x00");
    print_python_bytes(p);
    Py_XDECREF(p);

    p = PyBytes_FromString("What does the 'b' character do in front of a string literal?");
    print_python_bytes(p);
    Py_XDECREF(p);

    p = PyList_New(2);
    PyList_SetItem(p, 0, PyBytes_FromString("Hello"));
    PyList_SetItem(p, 1, PyBytes_FromString("World"));
    print_python_list(p);
    Py_XDECREF(p);

    return 0;
}
