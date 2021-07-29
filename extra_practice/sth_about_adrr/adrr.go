package sth_about_adrr

type People struct {
	Name string
	Age  int
}

type Worker struct {
	//工位
	Location int
	Sb       *People
}

func TryAdrr(worker *Worker) {
	worker.Sb = &People{
		"pengyuan", 18,
	}
	worker.Sb.Name = "yuanpeng"
	worker.Sb.Age = 19
	worker.Location = 11413
	return
}
