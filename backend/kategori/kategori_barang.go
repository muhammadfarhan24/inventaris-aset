package kategori

import (
	"database/sql"
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Kategori struct {
	ID   int    `json:"id"`
	Kode string `json:"kode_kategori"`
	Nama string `json:"nama_kategori"`
}

func connectDB() (*sql.DB, error) {
	return sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/inventaris_barang")
}

func GetKategori(w http.ResponseWriter, r *http.Request) {
	db, err := connectDB()
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, kode_kategori, nama_kategori FROM kategori_barang")
	if err != nil {
		http.Error(w, "Query error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var list []Kategori
	for rows.Next() {
		var k Kategori
		if err := rows.Scan(&k.ID, &k.Kode, &k.Nama); err == nil {
			list = append(list, k)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(list)
}

func TambahKategori(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var k Kategori
	err := json.NewDecoder(r.Body).Decode(&k)
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

	_, err = db.Exec("INSERT INTO kategori_barang (kode_kategori, nama_kategori) VALUES (?, ?)", k.Kode, k.Nama)
	if err != nil {
		http.Error(w, "Insert error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Kategori berhasil ditambahkan"})
}

func EditKategori(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]

	var k Kategori
	err := json.NewDecoder(r.Body).Decode(&k)
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

	_, err = db.Exec("UPDATE kategori_barang SET kode_kategori=?, nama_kategori=? WHERE id=?", k.Kode, k.Nama, id)
	if err != nil {
		http.Error(w, "Update error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Kategori berhasil diupdate"})
}

func HapusKategori(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	db, err := connectDB()
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM kategori_barang WHERE id = ?", id)
	if err != nil {
		http.Error(w, "Gagal menghapus", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
