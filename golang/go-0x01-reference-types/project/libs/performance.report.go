package libs

import (
	"fmt"
)

func NewStudent(student Student, grades SubjectGrades) *Student {
	return &Student{
		Name:         student.Name,
		RegNumber:    student.RegNumber,
		StudentClass: student.StudentClass,
		TotalMarks:   getTotalMarks(grades),
	}
}

func NewSubjectGrades(subjects SubjectGrades) *SubjectGrades {
	return &SubjectGrades{
		Math:          subjects.Math,
		English:       subjects.English,
		Kiswahili:     subjects.Kiswahili,
		Science:       subjects.Science,
		SocialStudies: subjects.SocialStudies,
	}
}

func getTotalMarks(subjects SubjectGrades) uint32 {
	return subjects.Math + subjects.English + subjects.Kiswahili + subjects.Science + subjects.SocialStudies
}

func NewExamReport() *ClassExamReport {
	return &ClassExamReport{}
}

func NewMySchoolReport() *MySchoolReport {
	return &MySchoolReport{}
}

func (report *MySchoolReport) PrintPerClass() {
	for _, classReport := range report.MySchoolPerformance {
		fmt.Printf("Class: %s <---> Number of Students: %d <---> Mean Grade: %.2f\n",
			classReport.Class, classReport.StudentList, classReport.MeanGrade)
	}
}

func (report *MySchoolReport) PrintAllReport() {
	for _, studentClass := range report.MySchoolPerformance {
		fmt.Printf("%s\n", studentClass.Class)
		//for _,
	}
}

func (std *Student) GetClass() string {
	return ""
}

func (std *Student) GetName() string {
	return ""
}

func (std *Student) GetMeanGrade(studentName string) string {
	return ""
}
