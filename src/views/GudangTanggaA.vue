<template>
  <div class="gudang-tangga-a">
    <div class="header-bar">
      <h2>Daftar Barang - Gudang Tangga A</h2>
      <button class="btn-tambah" @click="tambahBarang">+ Tambah Barang</button>
    </div>

    <div class="search-bar">
      <input type="text" v-model="searchQuery" placeholder="Cari Barang..." />
    </div>

    <table class="barang-table">
      <thead>
        <tr>
          <th>No</th>
          <th>Nama Barang</th>
          <th>Kategori</th>
          <th>Merk</th>
          <th>Status</th>
          <th>Deskripsi</th>
          <th>Aksi</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(barang, index) in filteredBarangList" :key="barang.id">
          <td>{{ index + 1 }}</td>
          <td>{{ barang.nama }}</td>
          <td>{{ barang.kategori_id }}</td>
          <td>{{ barang.merk_id }}</td>
          <td>{{ barang.status }}</td>
          <td>{{ barang.deskripsi }}</td>
          <td>
            <button @click="editBarang(barang.id)" class="btn-edit">Edit</button>
            <button @click="deleteBarang(barang.id)" class="btn-delete">Delete</button>
          </td>
        </tr>
        <tr v-if="filteredBarangList.length === 0">
          <td colspan="7" style="text-align: center;">Data Tidak Ditemukan</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>


<script>
export default {
    name: 'GudangTanggaA',
    data() {
        return {
            searchQuery: '',
            barangList: []
        }
    },
    computed: {
        filteredBarangList() {
            const query = this.searchQuery.toLowerCase();
            return this.barangList.filter(item =>
                item.nama.toLowerCase().includes(query) ||
                item.kategori_id.toLowerCase().includes(query)
            );
        }
    },
    methods: {
        fetchBarang() {
            fetch('http://localhost:3000/barang')
                .then(res => res.json())
                .then(data => {
                    this.barangList = data.filter(item => item.ruangan_id == 3); // hanya Gudang Imam A
                })
                .catch(err => console.error(err));
        },
        tambahBarang() {
            alert('Navigasi ke form tambah barang');
        },
        editBarang(id) {
            alert(`Edit barang dengan ID: ${id}`);
        },
        deleteBarang(id) {
            if (confirm('Yakin ingin menghapus barang ini')) {
                this.barangList = this.barangList.filter(item => item.id !== id);
            }
        }
    },
    mounted() {
        this.fetchBarang();
    }
}
</script>


<style scoped>
.gudang-tangga-a {
  padding: 20px;
  background-color: white;
  border-radius: 8px;
  box-shadow: 0 2px 5px rgba(0,0,0,0.1);
}

.gudang-tangga-a h2 {
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

.barang-table {
  width: 100%;
  border-collapse: collapse;
}

.barang-table th,
.barang-table td {
  padding: 12px;
  border-bottom: 1px solid #ccc;
  text-align: left;
}

.barang-table th {
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