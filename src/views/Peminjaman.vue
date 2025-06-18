<template>
  <div class="peminjaman">

    <!-- === DAFTAR PEMINJAMAN === -->
    <div class="header-bar">
      <h2>Daftar Peminjaman Aktif</h2>
      <button class="btn-tambah" @click="tambahPeminjam">+ Tambah Peminjam</button>
    </div>

    <!-- FORM TAMBAH PEMINJAMAN -->
    <div v-if="showForm" class="form-card">
      <div class="form-group">
        <label>Nama Barang</label>
        <select v-model.number="formData.barang_id">
          <option value="" disabled>Pilih Barang</option>
          <option v-for="barang in barangTersedia" :key="barang.id" :value="barang.id">
            {{ barang.nama }}
          </option>
        </select>
      </div>
      <div class="form-group">
        <label>Nama Peminjam</label>
        <input type="text" v-model="formData.peminjam" placeholder="Masukkan nama peminjam" />
      </div>
      <div class="form-group">
        <label>Tanggal Pinjam</label>
        <input type="date" v-model="formData.tanggal_peminjam" />
      </div>
      <div class="form-group">
        <label>Tanggal Kembali</label>
        <input type="date" v-model="formData.tanggal_kembali" />
      </div>
      <div class="form-group">
        <label>Status Pengembalian</label>
        <select v-model="formData.status_pengembalian">
          <option value="Dipinjam">Dipinjam</option>
          <option value="Dikembalikan">Dikembalikan</option>
        </select>
      </div>
      <div class="form-actions">
        <button @click="submitPeminjaman" class="btn-tambah">Simpan</button>
        <button @click="showForm = false" class="btn-delete">Batal</button>
      </div>
    </div>


    <div class="search-bar">
      <input type="text" v-model="searchQuery" placeholder="Cari peminjam atau barang..." />
    </div>

    <!-- DESKTOP TABLE -->
    <div class="table-container desktop-only">
      <table class="peminjaman-table">
        <thead>
          <tr>
            <th>ID</th>
            <th>Nama Peminjam</th>
            <th>Barang</th>
            <th>Tanggal Pinjam</th>
            <th>Tanggal Kembali</th>
            <th>Status</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(item, index) in filteredPeminjaman" :key="item.id">
            <!-- <td>{{ index + 1 }}</td> -->
            <td>{{ item.id }}</td>
            <td>{{ item.peminjam }}</td>
            <td>{{ item.nama_barang }}</td>
            <!-- <td>{{ item.merk }}</td>
            <td>{{ item.deskripsi }}</td> -->
            <td>{{ item.tanggal_peminjam }}</td>
            <!-- <td>{{ item.jatuhTempo }}</td> -->
            <td>{{ item.tanggal_kembali || '-' }}</td>
            <td>{{ item.status_pengembalian }}</td>
            <td>
              <button @click="editPeminjaman(item.id)" class="btn-edit">Edit</button>
              <button @click="hapusPeminjaman(item.id)" class="btn-delete">Hapus</button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- MOBILE CARD - Aktif -->
    <div class="mobile-card-container mobile-only">
      <div class="peminjaman-card" v-for="(item, index) in filteredPeminjaman" :key="item.id">
        <!-- <p><strong>No:</strong> {{ index + 1 }}</p> -->
        <p><strong>ID:</strong> {{ item.id }}</p>
        <p><strong>Nama:</strong> {{ item.peminjam }}</p>
        <p><strong>Nama Barang:</strong> {{ item.nama_barang }}</p>
        <!-- <p><strong>Merk:</strong> {{ item.merk }}</p>
        <p><strong>Deskripsi:</strong> {{ item.deskripsi }}</p> -->
        <p><strong>Tanggal Pinjam:</strong> {{ item.tanggal_peminjam }}</p>
        <!-- <p><strong>Jatuh Tempo:</strong> {{ item.jatuhTempo }}</p> -->
        <p><strong>Tanggal Kembali:</strong> {{ item.tanggal_kembali || '-' }}</p>
        <p><strong>Status:</strong> {{ item.status_pengembalian }}</p>
        <div class="card-actions">
          <button @click="editPeminjaman(item.id)" class="btn-edit">Edit</button>
          <button @click="hapusPeminjaman(item.id)" class="btn-delete">Hapus</button>
        </div>
      </div>
    </div>

    <!-- === RIWAYAT === -->
    <h2 class="section-title">Riwayat Peminjaman</h2>

    <!-- DESKTOP TABLE -->
    <div class="table-container desktop-only">
      <table class="peminjaman-table">
        <thead>
          <tr>
            <th>ID</th>
            <th>Nama Peminjam</th>
            <th>Barang</th>
            <th>Tanggal Pinjam</th>
            <th>Tanggal Kembali</th>
            <th>Status</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(riwayat, index) in riwayatPeminjaman" :key="riwayat.id">
            <!-- <td>{{ index + 1 }}</td> -->
            <td>{{ riwayat.id }}</td>
            <td>{{ riwayat.peminjam }}</td>
            <td>{{ riwayat.nama_barang }}</td>
            <!-- <td>{{ riwayat.deskripsi }}</td> -->
            <td>{{ riwayat.tanggal_peminjam }}</td>
            <td>{{ riwayat.tanggal_kembali }}</td>
            <!-- <td>{{ riwayat.tglKembali }}</td> -->
            <td>{{ riwayat.status_pengembalian }}</td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- MOBILE CARD - Riwayat -->
    <div class="mobile-card-container mobile-only riwayat-scroll">
      <div class="peminjaman-card" v-for="(riwayat, index) in riwayatPeminjaman" :key="riwayat.id">
        <!-- <p><strong>No:</strong> {{ index + 1 }}</p> -->
        <p><strong>ID:</strong> {{ riwayat.id }}</p>
        <p><strong>Nama:</strong> {{ riwayat.peminjam }}</p>
        <p><strong>Nama Barang:</strong> {{ riwayat.nama_barang }}</p>
        <!-- <p><strong>Deskripsi:</strong> {{ riwayat.deskripsi }}</p> -->
        <p><strong>Tanggal Pinjam:</strong> {{ riwayat.tanggal_peminjam }}</p>
        <!-- <p><strong>Jatuh Tempo:</strong> {{ riwayat.jatuhTempo }}</p> -->
        <p><strong>Tanggal Kembali:</strong> {{ riwayat.tanggal_kembali }}</p>
        <p><strong>Status:</strong> {{ riwayat.status_pengembalian }}</p>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      searchQuery: '',
      daftarPeminjaman: [],
      riwayatPeminjaman: [],
      barangTersedia: [],
      formData: {
        id: null,
        barang_id: '',
        peminjam: '',
        tanggal_peminjam: '',
        tanggal_kembali: '',
        status_pengembalian: 'Dipinjam'
      },
      showForm: false,
      isEdit: false
    };
  },
  computed: {
    filteredPeminjaman() {
      const q = this.searchQuery.toLowerCase();
      return this.daftarPeminjaman.filter(item =>
        item.peminjam.toLowerCase().includes(q) ||
        item.nama_barang.toLowerCase().includes(q)
      );
    }
  },
  methods: {
    async fetchPeminjaman() {
      try {
        const res = await fetch('http://localhost:3000/peminjaman');
        const data = await res.json();
        this.daftarPeminjaman = data.filter(item => item.status_pengembalian === 'Dipinjam');
        this.riwayatPeminjaman = data.filter(item => item.status_pengembalian === 'Dikembalikan');
      } catch (err) {
        console.error('Gagal mengambil data:', err);
      }
    },
    async fetchBarangTersedia() {
      try {
        const res = await fetch('http://localhost:3000/barang');
        const data = await res.json();
        this.barangTersedia = data.filter(b => b.status === "Tersedia");
      } catch (err) {
        console.error("Gagal mengambil data barang:", err);
      }
    },
    async tambahPeminjam() {
      this.showForm = !this.showForm;
      if (this.showForm && !this.isEdit) {
        await this.fetchBarangTersedia();
      }
    },
  async submitPeminjaman() {
  try {
    console.log('Barang ID:', this.formData.barang_id, typeof this.formData.barang_id);

    // Validasi manual
    if (
      this.formData.barang_id === null ||
      this.formData.barang_id === '' ||
      isNaN(this.formData.barang_id)
    ) {
      alert('Silakan pilih barang terlebih dahulu!');
      return;
    }

    const method = this.isEdit ? 'PUT' : 'POST';
    const url = this.isEdit
      ? `http://localhost:3000/peminjaman?id=${this.formData.id}`
      : 'http://localhost:3000/peminjaman';

    const res = await fetch(url, {
      method,
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(this.formData)
    });

    if (!res.ok) throw new Error('Gagal menyimpan');

    // Jangan pakai res.json() kalau POST karena backend tidak kirim JSON!
    if (this.isEdit) {
      const result = await res.json(); // PUT punya response JSON
      alert(result.message || 'Data berhasil diperbarui');
    } else {
      alert('Data berhasil ditambahkan');
    }

    this.showForm = false;
    this.resetFormData();
    this.fetchPeminjaman();
    this.fetchBarangTersedia(); // Refresh data barang setelah peminjaman

  } catch (err) {
    alert('Gagal menyimpan data');
    console.error('Error submit:', err);
  }
},
    editPeminjaman(id) {
      const item = [...this.daftarPeminjaman, ...this.riwayatPeminjaman].find(p => p.id === id);
      if (item) {
        this.formData = { ...item, barang_id: item.barang_id.toString() };
        this.isEdit = true;
        this.showForm = true;
      }
    },
    resetFormData() {
  this.formData = {
    id: null,
    barang_id: '',
    peminjam: '',
    tanggal_peminjam: '',
    tanggal_kembali: '',
    status_pengembalian: 'Dipinjam'
  };
  this.isEdit = false;
},
    
    async hapusPeminjaman(id) {
      if (confirm('Yakin ingin menghapus data ini?')) {
        try {
          const res = await fetch(`http://localhost:3000/peminjaman?id=${id}`, {
            method: 'DELETE'
          });
          const result = await res.json();
          alert(result.message);
          this.fetchPeminjaman();
        } catch (err) {
          alert('Gagal menghapus data');
        }
      }
    }
  },
  mounted() {
    this.fetchPeminjaman();
    this.fetchBarangTersedia();
    console.log('Barang tersedia:', this.barangTersedia);
  }
};
</script>



<style scoped>
.peminjaman {
  max-height: 80vh;
  overflow-y: auto;
  padding: 20px;
  background: white;
  border-radius: 8px;
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
}

.header-bar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 10px;
}

.btn-tambah {
  background-color: #27ae60;
  color: white;
  border: none;
  padding: 10px 16px;
  border-radius: 4px;
  cursor: pointer;
  font-weight: bold;
  transition: background-color 0.3s ease;
}

.btn-tambah:hover {
  background-color: #219150;
}

/* .section-title {
  font-size: 1.2rem;
  margin: 30px 0 15px;
  padding-bottom: 5px;
  border-bottom: 2px solid #eee;
} */

.search-bar {
  margin-bottom: 16px;
}

.search-bar input {
  width: 100%;
  padding: 10px;
  border: 1px solid #ccc;
  border-radius: 6px;
}

.table-container {
  overflow-x: auto;
  margin-bottom: 20px;
}

.peminjaman-table {
  width: 100%;
  border-collapse: collapse;
}

.peminjaman-table th,
.peminjaman-table td {
  padding: 10px;
  border-bottom: 1px solid #ddd;
  text-align: left;
}

.peminjaman-table th {
  background-color: #2c3e50;
  color: white;
}

.btn-edit,
.btn-delete {
  padding: 6px 10px;
  border-radius: 4px;
  border: none;
  color: white;
  cursor: pointer;
}

.btn-edit {
  background-color: #3498db;
}

.btn-edit:hover {
  background-color: #2980b9;
}

.btn-delete {
  background-color: #e74c3c;
}

.btn-delete:hover {
  background-color: #c0392b;
}

.form-card {
  background: #f9f9f9;
  border: 1px solid #ddd;
  padding: 16px;
  margin-bottom: 20px;
  border-radius: 8px;
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.05);
}

.form-group {
  display: flex;
  flex-direction: column;
  margin-bottom: 12px;
}

.form-group label {
  font-weight: bold;
  margin-bottom: 4px;
}

.form-group input,
.form-group select {
  padding: 8px;
  border: 1px solid #ccc;
  border-radius: 6px;
}

.form-actions {
  display: flex;
  gap: 10px;
  margin-top: 10px;
}


.mobile-only {
  display: none;
}

.desktop-only {
  display: block;
}

@media (max-width: 900px) {
  .desktop-only {
    display: none;
  }

  .mobile-only {
    display: flex;
    flex-direction: column;
  }

  .mobile-card-container {
    max-height: 400px;
    overflow-y: auto;
    margin-bottom: 30px;
  }

  .peminjaman-card {
    background: #f8f9fa;
    border-radius: 8px;
    padding: 14px;
    margin-bottom: 12px;
    box-shadow: 0 2px 5px rgba(0, 0, 0, 0.07);
  }
}
</style>
