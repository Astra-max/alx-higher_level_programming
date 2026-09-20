package libs

type SubjectGrades struct {
	Math          uint32
	English       uint32
	Kiswahili     uint32
	Science       uint32
	SocialStudies uint32
}

type Student struct {
	Name          string
	RegNumber     uint32
	StudentClass  string
	TotalMarks    uint32
	SubjectScores map[int]SubjectGrades
}

type ClassExamReport struct {
	Class       string
	MeanGrade   float32
	StudentList []Student
}

type MySchoolReport struct {
	MySchoolPerformance []ClassExamReport
}
