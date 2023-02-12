      <script>
    function a() {
        console.log("hello");
    }
</script>

<body class="bg-primary"></body>
<div class="navbar bg-base-100 bg-accent">
    <a class="btn btn-ghost btn-sm rounded-btn bg-secondary" href="http://localhost:5173">Auxili</a>
    <h1>--</h1>
    <h2><a class="btn btn-sm btn-ghost rounded-btn bg-secondary" href="http://localhost:5173/register">Register</a></h2>
    <h1>--</h1>
    <h2><a class="btn btn-ghost btn-sm rounded-btn bg-secondary" href="http://localhost:5173/login">Login</a></h2>
    
</div>

<div class="flex h-screen justify-center items-center">
    <div data-theme="mycodecompiled" class="card w-96 bg-secondary text-neutral-content">
        <div class="card-body items-center text-center">
            <h2 class="card-title text-neutral">Welcome to Auxili</h2>
            <br>
            <p class="card-subtitle text-primary">An app for vendors and customers; with emphasis on flexibility. This app uses golang for its backend web servers utilizing gin gonic, as well as rabbitmq for message broking.</p>
            <br>
            <br>
            <div class="card-actions">
                <a class="btn btn-outline btn-accent" href= "/login">Login</a>
                <a class="btn btn-outline btn-accent" href= "/register">Register</a>
                <a class="btn btn-outline btn-accent" href="/about>">About</a>
            </div>
        </div>
    </div>
</div>
