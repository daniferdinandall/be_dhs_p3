package namapackage

// func TestInsertDosen1(t *testing.T) {
// 	kode := "MYS"
// 	nama := "M. Yusril Helmi Setyawan, S.Kom.,"
// 	hp := "+62807117405"

// 	hasil := module.InsertDosen(kode, nama, hp)
// 	fmt.Println(hasil)
// }
// func TestInsertDosen2(t *testing.T) {
// 	kode := "RMA"
// 	nama := "Rolly Maulana Awangga,S.T.,MT.,CAIP, SFPC."
// 	hp := "+62810118609"

// 	hasil := module.InsertDosen(kode, nama, hp)
// 	fmt.Println(hasil)
// }
// func TestInsertDosen3(t *testing.T) {
// 	kode := "MNF"
// 	nama := "Mohamad Nurkamal Fauzan, S.T., M.T., SFPC"
// 	hp := "+62802058005"

// 	hasil := module.InsertDosen(kode, nama, hp)
// 	fmt.Println(hasil)
// }
// func TestInsertDosen4(t *testing.T) {
// 	kode := "RHA"
// 	nama := "Roni Habibi, S.Kom., M.T., SFPC"
// 	hp := "+62823127804"

// 	hasil := module.InsertDosen(kode, nama, hp)
// 	fmt.Println(hasil)
// }
// func TestInsertDosen5(t *testing.T) {
// 	kode := "RAN"
// 	nama := "Roni Andarsyah, S.T., M.Kom., SFPC"
// 	hp := "+62820058801"

// 	hasil := module.InsertDosen(kode, nama, hp)
// 	fmt.Println(hasil)
// }
// func TestInsertDosen6(t *testing.T) {
// 	kode := "CPY"
// 	nama := "Cahyo Prianto, S.Pd., M.T.,CDSP, SFPC"
// 	hp := "+62827078401"

// 	hasil := module.InsertDosen(kode, nama, hp)
// 	fmt.Println(hasil)
// }
// func TestInsertDosen7(t *testing.T) {
// 	kode := "SFP"
// 	nama := " Syafrial Fachri Pane,ST. MTI,EBDP.CDSP,SFPC"
// 	hp := "+62815048901"

// 	hasil := module.InsertDosen(kode, nama, hp)
// 	fmt.Println(hasil)
// }
// func TestInsertDosen8(t *testing.T) {
// 	kode := "RNS"
// 	nama := "Rd. Nuraini Siti Fatonah, S.S., M.Hum., SFPC"
// 	hp := "+62815048901"

// 	hasil := module.InsertDosen(kode, nama, hp)
// 	fmt.Println(hasil)
// }
// func TestInsertDosen9(t *testing.T) {
// 	kode := "NHH"
// 	nama := "Nisa Hanum Harani, S.Kom., M.T.,CDSP, SFPC"
// 	hp := "+62815048901"

// 	hasil := module.InsertDosen(kode, nama, hp)
// 	fmt.Println(hasil)
// }
// func TestInsertDosen10(t *testing.T) {
// 	kode := "WIR"
// 	nama := "Woro Isti Rahayu, S.T., M.T., SFPC"
// 	hp := "+62815107901"

// 	hasil := module.InsertDosen(kode, nama, hp)
// 	fmt.Println(hasil)
// }
// func TestInsertDosen11(t *testing.T) {
// 	kode := "NRI"
// 	nama := "Noviana Riza, S.Si., M.T., SFPC"
// 	hp := "+62803117607"

// 	hasil := module.InsertDosen(kode, nama, hp)
// 	fmt.Println(hasil)
// }

// func TestInsertMatkul1(t *testing.T) {
// 	kode := "TI41122"
// 	nama := "ALGORITMA DAN STRUKTUR DATA 2"
// 	sks := 3
// 	dosen := model.Dosen{
// 		Nama: "Mohamad Nurkamal Fauzan, S.T., M.T., SFPC",
// 	}

// 	hasil := module.InsertMatkul(kode, nama, sks, dosen)
// 	fmt.Println(hasil)
// }
// func TestInsertMatkul2(t *testing.T) {
// 	kode := "TI41061"
// 	nama := "ALGORITMA DAN STRUKTUR DATA I"
// 	sks := 3
// 	dosen := model.Dosen{
// 		Nama: "Mohamad Nurkamal Fauzan, S.T., M.T., SFPC",
// 	}

// 	hasil := module.InsertMatkul(kode, nama, sks, dosen)
// 	fmt.Println(hasil)
// }
// func TestInsertMatkul3(t *testing.T) {
// 	kode := "TI41092"
// 	nama := "ALJABAR LINIER"
// 	sks := 2
// 	dosen := model.Dosen{
// 		Nama: "Cahyo Prianto, S.Pd., M.T.,CDSP, SFPC",
// 	}

// 	hasil := module.InsertMatkul(kode, nama, sks, dosen)
// 	fmt.Println(hasil)
// }
// func TestInsertMatkul4(t *testing.T) {
// 	kode := "PPI01040"
// 	nama := "BAHASA INGGRIS 1"
// 	sks := 2
// 	dosen := model.Dosen{
// 		Nama: "RRd. Nuraini Siti Fatonah, S.S., M.Hum., SFPC",
// 	}

// 	hasil := module.InsertMatkul(kode, nama, sks, dosen)
// 	fmt.Println(hasil)
// }
// func TestInsertMatkul5(t *testing.T) {
// 	kode := "PPI02060"
// 	nama := "BAHASA INGGRIS 2"
// 	sks := 2
// 	dosen := model.Dosen{
// 		Nama: "Rd. Nuraini Siti Fatonah, S.S., M.Hum., SFPC",
// 	}

// 	hasil := module.InsertMatkul(kode, nama, sks, dosen)
// 	fmt.Println(hasil)
// }
// func TestInsertMatkul6(t *testing.T) {
// 	kode := "TI41091"
// 	nama := "LITERASI TEKNOLOGI"
// 	sks := 3
// 	dosen := model.Dosen{
// 		Nama: "Roni Habibi, S.Kom., M.T., SFPC",
// 	}

// 	hasil := module.InsertMatkul(kode, nama, sks, dosen)
// 	fmt.Println(hasil)
// }
// func TestInsertMatkul7(t *testing.T) {
// 	kode := "TI43132"
// 	nama := "BASIS DATA I "
// 	sks := 4
// 	dosen := model.Dosen{
// 		Nama: "Syafrial Fachri Pane,ST. MTI,EBDP.CDSP,SFPC",
// 	}

// 	hasil := module.InsertMatkul(kode, nama, sks, dosen)
// 	fmt.Println(hasil)
// }
// func TestInsertMatkul8(t *testing.T) {
// 	kode := "TI41011"
// 	nama := "MATEMATIKA DISKRIT"
// 	sks := 2
// 	dosen := model.Dosen{
// 		Nama: "Cahyo Prianto, S.Pd., M.T.,CDSP, SFPC",
// 	}

// 	hasil := module.InsertMatkul(kode, nama, sks, dosen)
// 	fmt.Println(hasil)
// }
// func TestInsertMatkul9(t *testing.T) {
// 	kode := "TI41071"
// 	nama := "PEMOGRAMAN I"
// 	sks := 3
// 	dosen := model.Dosen{
// 		Nama: "Roni Andarsyah, S.T., M.Kom., SFPC",
// 	}

// 	hasil := module.InsertMatkul(kode, nama, sks, dosen)
// 	fmt.Println(hasil)
// }
// func TestInsertMatkul10(t *testing.T) {
// 	kode := "TI42092"
// 	nama := "BASIS DATA I "
// 	sks := 4
// 	dosen := model.Dosen{
// 		Nama: "Syafrial Fachri Pane,ST. MTI,EBDP.CDSP,SFPC",
// 	}

// 	hasil := module.InsertMatkul(kode, nama, sks, dosen)
// 	fmt.Println(hasil)
// }
