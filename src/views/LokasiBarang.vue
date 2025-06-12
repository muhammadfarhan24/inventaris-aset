<template>
  <div class="lokasi-barang">
    <div class="header-bar">
        <h2>Daftar Lokasi Barang</h2>
        <button class="btn-tambah" @click="tambahLokasi">+ Tambah Lokasi</button>
    </div>

    <!-- Kolom Search -->
    <div class="search-bar">
        <input type="text" v-model="searchQuery" placeholder="Cari Lokasi..." />
    </div> 

    <table class="lokasi-table">
      <thead>
        <tr>
          <th>No</th>
          <th>Kode Lokasi</th>
          <th>Lokasi</th>
          <th>Aksi</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(lokasi, index) in lokasiList" :key="lokasi.id">
          <td>{{ index + 1 }}</td>
          <td>{{ lokasi.kode }}</td>
          <td>{{ lokasi.nama }}</td>
          <td>
            <button @click="editLokasi(lokasi.id)" class="btn-edit">Edit</button>
            <button @click="deleteLokasi(lokasi.id)" class="btn-delete">Delete</button>
          </td>
        </tr>

        <tr v-if="filteredLokasiList.length === 0">
            <td colspan="4" style="text-align: center;">Data Tidak Ditemukan</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
export default {
  name: 'LokasiBarang',
  data() {
    return {
      searchQuery: '',  
      lokasiList: [
        { id: 1, kode: 'LOK001', nama: 'Masjid' },
      ]
    }
  },
  computed: {
    filteredLokasiList() {
        const query = this.searchQuery.toLowerCase();
        return this.lokasiList.filter(kat =>
            kat.nama.toLowerCase().includes(query) || kat.kode.toLowerCase().includes(query)
        );
    }
  },
  methods: {
    editLokasi(id) {
      alert(`Edit lokasi dengan ID: ${id}`);
      // Navigasi ke form edit atau tampilkan modal edit
    },
    deleteLokasi(id) {
      if (confirm('Yakin ingin menghapus lokasi ini?')) {
        this.lokasiList = this.lokasiList.filter(kat => kat.id !== id);
      }
    }
  }
}
</script>

<style scoped>
.lokasi-barang {
  padding: 20px;
  background-color: white;
  border-radius: 8px;
  box-shadow: 0 2px 5px rgba(0,0,0,0.1);
}

.lokasi-barang h2 {
  margin-bottom: 20px;
}

.header-bar {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 20px;
}

.btn-tambah {
    background-color: #27ae60;
    color: white;
    border: none;
    padding: 10px 16px;
    border-radius: 4px;
    cursor: pointer;
    font-weight: bold;
}

.btn-tambah:hover {
    background-color: #219150;
}

.search-bar {
    margin-bottom: 20px;
}

.search-bar input {
    width: 100%;
    padding: 10px 12px;
    font-size: 1rem;
    border: 1px solid #ccc;
    border-radius: 4px;
}

.lokasi-table {
  width: 100%;
  border-collapse: collapse;
}

.lokasi-table th,
.lokasi-table td {
  padding: 12px;
  border-bottom: 1px solid #ccc;
  text-align: left;
}

.lokasi-table th {
  background-color: #2c3e50;
  color: white;
}

.btn-edit {
  background-color: #3498db;
  color: white;
  border: none;
  padding: 6px 10px;
  border-radius: 4px;
  margin-right: 5px;
  cursor: pointer;
}

.btn-delete {
  background-color: #e74c3c;
  color: white;
  border: none;
  padding: 6px 10px;
  border-radius: 4px;
  cursor: pointer;
}

.btn-edit:hover {
  background-color: #2980b9;
}

.btn-delete:hover {
  background-color: #c0392b;
}
</style>
