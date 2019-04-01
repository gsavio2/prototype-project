<template>
    <footer>
        <div class="footer-app">
            <div class="footer">
                <router-link class="mx-auto" to="/"> BREJARIA </router-link> &copy; 2019 – {{ new Date().getFullYear() }}   
            </div>        
        </div>
    </footer>
</template>

<script>
export default {
    name: 'footer-app'
}
</script>

<style scoped>
    .footer{
        border-top: 1px solid slategray;
        margin: 5px;
        padding: 10px;
    }
</style>
