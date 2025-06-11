package barang

import (
	"database/sql"
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Barang struct {
	ID         int    `json:"id"`
	Nama       string `json:"nama"`
	Kategori   string `json:"kategori_id"`
	Merk       string `json:"merk_id"`
	Status     string `json:"status"`
	Ruangan    string `json:"ruangan_id"`
	Deskripsi  string `json:"deskripsi"`
	Created_at string `json:"created_at"`
}

func connectDB() (*sql.DB, error) {
	return sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/inventaris_barang")
}

func GetBarang(w http.ResponseWriter, r *http.Request) {
	db, err := connectDB()
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, nama, kategori_id, merk_id, status, ruangan_id, deskripsi, created_at FROM barang")
	if err != nil {
		http.Error(w, "Query error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var list []Barang
	for rows.Next() {
		var b Barang
		err := rows.Scan(&b.ID, &b.Nama, &b.Kategori, &b.Merk, &b.Status, &b.Ruangan, &b.Deskripsi, &b.Created_at)
		if err != nil {
			continue
		}
		list = append(list, b)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(list)
}

func TambahBarang(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var b Barang
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	db, err := connectDB()
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, err = db.Exec(`
		INSERT INTO barang (nama, kategori_id, merk_id, status, ruangan_id, deskripsi)
		VALUES (?, ?, ?, ?, ?, ?)`,
		b.Nama, b.Kategori, b.Merk, b.Status, b.Ruangan, b.Deskripsi)

	if err != nil {
		http.Error(w, "Insert error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Barang berhasil ditambahkan"})
}
