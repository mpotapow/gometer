<template>
    <div>
        <nav class="navbar bg-gradient">
            <a class="navbar-brand" href="#">
                GoMeter
            </a>
            <form class="form-inline">
                <!-- <input class="form-control mr-sm-2" type="search" placeholder="Search" aria-label="Search"> -->
                <!-- <button class="btn btn-outline-success my-2 my-sm-0" type="submit">Search</button> -->
                <div>
                    <div class="dropdown">
                    <button class="btn btn-secondary dropdown-toggle" type="button" id="dropdownMenu2" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
                        Dropdown
                    </button>
                    <div class="dropdown-menu" aria-labelledby="dropdownMenu2">
                        <button class="dropdown-item" type="button">Action</button>
                        <button class="dropdown-item" type="button">Another action</button>
                        <button class="dropdown-item" type="button">Something else here</button>
                    </div>
                    </div>
                </div>
            </form>
        </nav>
        <div class="container">
            <router-view></router-view>
        </div>
    </div>
</template>

<script>
    export default {
        name: 'main-layout',
        components: {

        }
    }
</script>
<style>
    .bg-gradient {
        background: rgb(232,73,189);
        background: linear-gradient(90deg, rgba(232,73,189,1) 0%, rgba(253,29,240,1) 50%, rgba(171,56,219,1) 100%);
    }
</style>