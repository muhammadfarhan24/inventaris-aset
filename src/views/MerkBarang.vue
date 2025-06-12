<template>
  <div class="merk-barang">
    <div class="header-bar">
        <h2>Daftar Merk Barang</h2>
        <button class="btn-tambah" @click="tambahMerk">+ Tambah Merk</button>
    </div>

    <!-- Kolom Search -->
    <div class="search-bar">
        <input type="text" v-model="searchQuery" placeholder="Cari Merk..." />
    </div> 

    <table class="merk-table">
      <thead>
        <tr>
          <th>No</th>
          <th>Kode Merk</th>
          <th>Nama Merk</th>
          <th>Aksi</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(merk, index) in filteredMerkList" :key="merk.id">
          <td>{{ index + 1 }}</td>
          <td>{{ merk.kode_merk }}</td>
          <td>{{ merk.nama_merk }}</td>
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
      merkList: []
    };
  },
  computed: {
    filteredMerkList() {
      const query = this.searchQuery.toLowerCase();
      return this.merkList.filter(m =>
        (m.nama_merk || '').toLowerCase().includes(query) ||
        (m.kode_merk || '').toLowerCase().includes(query)
      );
    }
  },
  methods: {
    fetchMerk() {
      fetch('http://localhost:3000/merk')
        .then(res => res.json())
        .then(data => {
          this.merkList = data;
        })
        .catch(err => {
          console.error('Gagal ambil data merk:', err);
        });
    },
    tambahMerk() {
      const kode = prompt('Masukkan kode merk:');
      const nama = prompt('Masukkan nama merk:');
      if (!kode || !nama) return;

      fetch('http://localhost:3000/merk', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ kode_merk: kode, nama_merk: nama })
      })
        .then(res => res.json())
        .then(() => this.fetchMerk())
        .catch(err => {
          console.error('Gagal tambah merk:', err);
        });
    },
    deleteMerk(id) {
      if (!confirm('Yakin ingin menghapus merk ini?')) return;

      fetch(`http://localhost:3000/merk/${id}`, {
        method: 'DELETE'
      })
        .then(() => this.fetchMerk())
        .catch(err => {
          console.error('Gagal hapus merk:', err);
        });
    },
    editMerk(id) {
      alert(`Edit merk dengan ID: ${id}`);
      // fitur edit bisa dikembangkan nanti
    }
  },
  mounted() {
    this.fetchMerk();
  }
};
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
