package main

type Matrix struct {
	Data []int
	N    int
	M    int
}

func (m *Matrix) Get(i, j int) int {
	return m.Data[i*m.N+j]
}

func (m *Matrix) Set(i, j, val int) {
	m.Data[i*m.N+j] = val
}

func (m *Matrix) Transpose() {
	m.M, m.N = m.N, m.M
}
