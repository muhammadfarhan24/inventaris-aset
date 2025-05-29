<template>
    <div class="login-page">
        <div class="login-card card shadow-sm p-4">
            <h4 class="text-center mb-3">Login admin</h4>
            <form @submit.prevent="handleLogin">
                <div class="mb-3">
                    <label for="email" class="form-label">email</label>
                    <input v-model="email" type="email" id="email" class="form-control" required />
                </div>
                <div class="mb-3">
                    <label for="password" class="form-label">Password</label>
                    <input v-model="password" type="password" id="email" class="form-control" required />
                </div>
                <button type="submit" class="btn btn-primary w-100">Login</button>
            </form>
        </div>
    </div>
</template>

<script>
export default {
    name: 'LoginPage',
    data() {
        return {
            email: '', 
            password: ''
        }
    },
    methods: {
        handleLogin() {
            if (this.email === 'admin@example.com' && this.password === 'admin') {
                this.$router.push('/dashboard')
            } else {
                alert('Email atau Password salah.')
            }
        }
    }
}
</script>

<style scoped>
.login-page {
    min-height: 100vh;
    display: flex;
    justify-content: center;
    align-items: center;
    background-color: #f8f9fa;
    padding: 1rem;
}

.login-card {
    width: 100%;
    max-width: 320px;
    background-color: white;
    border-radius: 8px;
}

</style>