<template>
    <div class="gudang-lantai-2">
        <div class="header-bar">
            <h2>Daftar Barang - Gudang Lantai 2</h2>
            <button class="btn-tambah" @click="tambahBarang">+ Tambah Barang</button>
        </div>

        <!-- Kolom Search -->
        <div class="search-bar">
            <input type="text" v-model="searchQuery" placeholder="Cari Barang..." />
        </div>

        <table class="barang-table">
            <thead>
                <tr>
                    <th>No</th>
                    <th>Kode Barang</th>
                    <th>Nama Barang</th>
                    <th>Kategori</th>
                    <th>Jumlah</th>
                    <th>Aksi</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="(barang, index) in filteredBarangList" :key="barang.id">
                    <td>{{ index + 1 }}</td>
                    <td>{{ barang.kodeBarang }}</td>
                    <td>{{ barang.namaBarang }}</td>
                    <td>{{ barang.kategori }}</td>
                    <td>{{ barang.jumlah }}</td>
                    <td>
                        <button @click="editBarang(barang.id)" class="btn-edit">Edit</button>
                        <button @click="deleteBarang(barang.id)" class="btn-delete">Delete</button>
                    </td>
                </tr>
                <tr v-if="filteredBarangList.length === 0">
                    <td colspan="6" style="text-align: center;">Data Tidak Ditemukan</td>
                </tr>
            </tbody>
        </table>
    </div>
</template>

<script>
export default {
    name: 'GudangLantai2',
    data() {
        return {
            searchQuery: '',
            barangList: [
                { id: 1, kodeBarang: 'BRG001', namaBarang: 'Kursi Lipat', kategori: 'Perabot', jumlah: 12 },
                { id: 2, kodeBarang: 'BRG002', namaBarang: 'Meja Lipat', kategori: 'Perabot', jumlah: 5 }
            ]
        }
    },
    computed: {
        filteredBarangList() {
            const query = this.searchQuery.toLocaleLowerCase();
            return this.barangList.filter(item =>
                item.namaBarang.toLocaleLowerCase().includes(query) ||
                item.kodeBarang.toLocaleLowerCase().includes(query) ||
                item.kategori.toLocaleLowerCase().includes(query)
            );
        }
    },
    methods: {
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
    }
}
</script>

<style scoped>
.gudang-lantai-2 {
  padding: 20px;
  background-color: white;
  border-radius: 8px;
  box-shadow: 0 2px 5px rgba(0,0,0,0.1);
}

.gudang-lantai-2 h2 {
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