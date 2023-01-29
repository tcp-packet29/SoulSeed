<script>

    const axios = require("axios");

    function createUser() {
        
    }

</script>

<body class="bg-primary"></body>

<div class="flex h-screen justify-center items-center">
    <div data-theme="mycodecompiled" class="card w-96 bg-secondary bg-base-200 text-neutral-content shadow-xl">
        <div class="card-body items-center text-center">
            <h2 class="card-title text-neutral">Register</h2>
            <p class="card-subtitle text-primary">An app for vendors and customers; with emphasis on flexibility.</p>
            <input type="text" placeholder="Username" class="input input-bordered input-accent w-full max-w-xs" />
            <input type="text" placeholder="Password" class="input input-bordered input-accent w-full max-w-xs" />
            <div class="card-actions">
                <button class="btn btn-outline btn-accent" on:click={createUser}>Register</button>
            </div>
        </div>
    </div>
</div>