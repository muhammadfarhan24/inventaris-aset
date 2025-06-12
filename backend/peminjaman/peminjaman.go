package peminjaman

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type Peminjaman struct {
	ID                 int    `json:"id"`
	NamaBarang         string `json:"nama_barang"`
	Peminjam           string `json:"peminjam"`
	TanggalPeminjam    string `json:"tanggal_peminjam"`
	TanggalKembali     string `json:"tanggal_kembali"`
	StatusPengembalian string `json:"status_pengembalian"`
}

func connectDB() (*sql.DB, error) {
	return sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/inventaris_barang")
}

func GetAllPeminjaman(w http.ResponseWriter, r *http.Request) {
	db, err := connectDB()
	if err != nil {
		http.Error(w, "DB connection error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query(`
	SELECT id, nama_barang, peminjam, tanggal_peminjam, tanggal_kembali, status_pengembalian
	FROM peminjaman
`)

	if err != nil {
		http.Error(w, "Query error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var list []Peminjaman
	for rows.Next() {
		var p Peminjaman
		if err := rows.Scan(&p.ID, &p.NamaBarang, &p.Peminjam, &p.TanggalPeminjam, &p.TanggalKembali, &p.StatusPengembalian); err != nil {
			http.Error(w, "Scan error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		list = append(list, p)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(list)
}

func TambahPeminjaman(w http.ResponseWriter, r *http.Request) {
	var data Peminjaman
	json.NewDecoder(r.Body).Decode(&data)

	db, err := connectDB()
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, err = db.Exec(`
		INSERT INTO peminjaman (nama_barang, peminjam, tanggal_peminjam, tanggal_kembali, status_pengembalian)
		VALUES (?, ?, ?, ?, ?)`,
		data.NamaBarang, data.Peminjam, data.TanggalPeminjam, data.TanggalKembali, data.StatusPengembalian)

	if err != nil {
		http.Error(w, "Gagal tambah data: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func EditPeminjaman(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID tidak valid", http.StatusBadRequest)
		return
	}

	var data Peminjaman
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Data tidak valid: "+err.Error(), http.StatusBadRequest)
		return
	}

	db, err := connectDB()
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, err = db.Exec(`
		UPDATE peminjaman
		SET nama_barang = ?, peminjam = ?, tanggal_peminjam = ?, tanggal_kembali = ?, status_pengembalian = ?
		WHERE id = ?`,
		data.NamaBarang, data.Peminjam, data.TanggalPeminjam, data.TanggalKembali, data.StatusPengembalian, id)

	if err != nil {
		http.Error(w, "Gagal update: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Data peminjaman berhasil diupdate",
	})
}

func HapusPeminjaman(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID tidak valid", http.StatusBadRequest)
		return
	}

	db, err := connectDB()
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	result, err := db.Exec("DELETE FROM peminjaman WHERE id = ?", id)
	if err != nil {
		http.Error(w, "Gagal hapus", http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "ID tidak ditemukan", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Data peminjaman berhasil dihapus",
	})
}
