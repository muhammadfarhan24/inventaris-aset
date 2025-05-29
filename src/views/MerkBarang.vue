<template>
  <div class="merk-barang">
    <div class="header-bar">
        <h2>Daftar Merk Barang</h2>
        <button class="btn-tambah" @click="tambahMerk">+ Tambah Merk</button>
    </div>

    <!-- Kolom Search -->
    <div class="seacrh-bar">
        <input type="text" v-model="searchQuery" placeholder="Cari Merk..." />
    </div> 

    <table class="merk-table">
      <thead>
        <tr>
          <th>No</th>
          <th>Kode Merk</th>
          <th>Lokasi</th>
          <th>Aksi</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(merk, index) in merkList" :key="merk.id">
          <td>{{ index + 1 }}</td>
          <td>{{ merk.kode }}</td>
          <td>{{ merk.nama }}</td>
          <td>
            <button @click="editMerk(merk.id)" class="btn-edit">Edit</button>
            <button @click="deleteMerk(merk.id)" class="btn-delete">Delete</button>
          </td>
        </tr>

        <tr v-if="filteredMerkList.length === 0">
            <td colspan="4" style="text-align: center;">Data Tidak Ditemukan</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
export default {
  name: 'MerkBarang',
  data() {
    return {
      searchQuery: '',  
      merkList: [
        { id: 1, kode: 'MRK001', nama: 'TOA' },
        { id: 2, kode: 'MRK002', nama: 'COBA' },
        { id: 3, kode: 'MRK003', nama: 'COBA' }
      ]
    }
  },
  computed: {
    filteredMerkList() {
        const query = this.searchQuery.toLowerCase();
        return this.merkList.filter(kat =>
            kat.nama.toLowerCase().includes(query) || kat.kode.toLowerCase().includes(query)
        );
    }
  },
  methods: {
    editMerk(id) {
      alert(`Edit merk dengan ID: ${id}`);
      // Navigasi ke form edit atau tampilkan modal edit
    },
    deleteMerk(id) {
      if (confirm('Yakin ingin menghapus lokasi ini?')) {
        this.merkList = this.merkList.filter(kat => kat.id !== id);
      }
    }
  }
}
</script>

<style scoped>
.merk-barang {
  padding: 20px;
  background-color: white;
  border-radius: 8px;
  box-shadow: 0 2px 5px rgba(0,0,0,0.1);
}

.merk-barang h2 {
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

.merk-table {
  width: 100%;
  border-collapse: collapse;
}

.merk-table th,
.merk-table td {
  padding: 12px;
  border-bottom: 1px solid #ccc;
  text-align: left;
}

.merk-table th {
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
