package ruangan

import (
	"database/sql"
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Ruangan struct {
	ID    int    `json:"id"`
	Nama  string `json:"nama"`
	Jenis string `json:"jenis"`
}

type Barang struct {
	ID        int    `json:"id"`
	Nama      string `json:"nama"`
	Kategori  string `json:"kategori_id"`
	Merk      string `json:"merk_id"`
	Status    string `json:"status"`
	Ruangan   string `json:"ruangan_id"`
	Deskripsi string `json:"deskripsi"`
	CreatedAt string `json:"created_at"`
}

func connectDB() (*sql.DB, error) {
	return sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/inventaris_barang")
}

// Ambil semua ruangan
func GetRuangan(w http.ResponseWriter, r *http.Request) {
	db, err := connectDB()
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, nama, jenis FROM ruangan")
	if err != nil {
		http.Error(w, "Query error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var list []Ruangan
	for rows.Next() {
		var r Ruangan
		if err := rows.Scan(&r.ID, &r.Nama, &r.Jenis); err == nil {
			list = append(list, r)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(list)
}

// Tambah ruangan
func TambahRuangan(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var rgn Ruangan
	if err := json.NewDecoder(r.Body).Decode(&rgn); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	db, err := connectDB()
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO ruangan (nama, jenis) VALUES (?, ?)", rgn.Nama, rgn.Jenis)
	if err != nil {
		http.Error(w, "Insert error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Ruangan berhasil ditambahkan"})
}

// Ambil barang berdasarkan ID ruangan
func BarangByRuangan(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing ruangan_id", http.StatusBadRequest)
		return
	}

	db, err := connectDB()
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query(`SELECT id, nama, kategori_id, merk_id, status, ruangan_id, deskripsi, created_at 
		FROM barang WHERE ruangan_id = ?`, id)
	if err != nil {
		http.Error(w, "Query error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var list []Barang
	for rows.Next() {
		var b Barang
		if err := rows.Scan(&b.ID, &b.Nama, &b.Kategori, &b.Merk, &b.Status, &b.Ruangan, &b.Deskripsi, &b.CreatedAt); err == nil {
			list = append(list, b)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(list)
}
