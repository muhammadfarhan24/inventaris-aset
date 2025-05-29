<template>
    <div class="dashboard">
        <main class="main-content">
            <header class="header">
                <button class="toggle-btn d-md-none" @click="sidebarOpen = !sidebarOpen">
                    <i class="bi bi-list" style="font-size: 1.5rem; color: #333;"></i>
                </button>
                <h1>Dashboard</h1>
            </header>
            <section class="content">
                <h2>Selamat datang di Dashboard Yayasan AL-Istiqomah</h2>
                <p>Berikut adalah ringkasan data inventaris</p>
                <div class="cards">
                    <button class="card" @click="activeView = 'barang'">
                        <h3>Total Barang</h3>
                        <p>120</p>
                    </button>
                    <button class="card" @click="activeView = 'kategori'">
                        <h3>Jumlah Kategori</h3>
                        <p>8</p>
                    </button>
                    <button class="card" @click="activeView = 'ruangan'">
                        <h3>Jumlah Ruangan</h3>
                        <p>7</p>
                    </button>
                    <button class="card" @click="activeView = 'pengguna'">
                        <h3>Total Pengguna</h3>
                        <p>5</p>
                    </button>
                </div>

                <div class="content-area" style="margin-top: 30px;">
                    <div v-if="activeView === 'barang'">
                        <h3>Daftar Barang</h3>
                        <p>Ini adalah daftar barang (bisa diganti dengan table)</p>
                    </div>
                    <div v-if="activeView === 'kategori'">
                        <h3>Daftar Kategori</h3>
                        <p>Ini adalah daftar kategori</p>
                    </div>
                    <div v-if="activeView === 'ruangan'">
                        <h3>Daftar Ruangan</h3>
                        <p>Ini adalah daftar ruangan</p>
                    </div>
                    <div v-if="activeView === 'pengguna'">
                        <h3>Daftar Pengguna</h3>
                        <p>Ini adalah daftar pengguna</p>
                    </div>
                </div>
            </section>
        </main>
    </div>
</template>

<script>
export default {
    name: 'DashboardPage',
    data() {
        return { activeView: '' }
    }
}
</script>

<style>
.dashboard {
    display: flex;
    height: 100vh;
    overflow: hidden;
}



.main-content {
    flex: 1;
    background-color: #f4f4f4;
    display: flex;
    flex-direction: column;
}

.header {
    background-color: white;
    padding: 20px;
    border-bottom: 1px solid #ddd;
}

.content {
    padding: 20px;
}

.toggle-btn {
    background: none;
    border: none;
    cursor: pointer;
    padding: 4px;
}

.cards {
    display: flex;
    flex-wrap: wrap;
    gap: 20px;
    margin-top: 20px;
}

.card:hover {
    background-color: #ecf0f1;
}

.card {
    background-color: white;
    padding: 20px;
    flex: 1 1 200px;
    border-radius: 8px;
    box-shadow: 0 2px 5px rgba( 0, 0, 0, 0.1);
    text-align: center;
    cursor: pointer;
    border: none;
    transition: background-color 0.2s;
}

.card h3 {
    margin-bottom: 10px;
    color: #2c3e50;
}

.card p {
    font-size: 24px;
    font-weight: bold;
    color: #2980b9;
}

.summary-table {
    width: 100%;
    border-collapse: collapse;
    margin-top: 20px;
    background-color: white;
    border-radius: 8px;
    overflow: hidden;
    box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
}

.summary-table th, 
.summary-table td {
    padding: 12px 16px;
    text-align: left;
    border-bottom: 1px solid #ddd;
}

.summary-table thead {
    background-color: #2c3e50;
    color: white;
}

.summary-table tbody tr:hover {
    background-color: #f1f1f1;
}

/* Responsive behavior */
@media (max-width: 768px) {
    .sidebar {
        position: absolute;
        left: 0;
        top: 0;
        transform: translateX(-100%);
        z-index: 999;
    }

    .sidebar.open {
        transform: translateX(0);
    }

    .main-content {
        padding-top: 10px;
    }
}
</style>