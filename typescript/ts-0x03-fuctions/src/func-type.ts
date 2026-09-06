type callMe = (n1: number, n2: number)=>number;

function helpMe(num1: number, num2: number, math: callMe): number {
    return math(num1, num2)
}

function add(n1: number, n2: number): number {
    return n1 + n2;
}

function mult(n1: number, n2: number): number {
    return n1 * n2;
}

function sub(n1: number, n2: number): number {
    return n1 - n2;
}

function mod(n1: number, n2: number): number {
    return n1 % n2;
}


console.log(helpMe(3,4, add))