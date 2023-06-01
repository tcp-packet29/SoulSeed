<script>
    import {page} from '$app/stores'; //sveltekit sdk is cooldepsite me wanting to try remix
    let stuff = {
        404: "The data you are trying to fetch does not exist.",
        500: "The server has encountered an error. Oops!",
        401: "You are not authorized to access this page.",
    }

    let jwt = stuff[$page.status]
</script>

<body class="bg-secondary h-screen">
<br>

<div class="divider"><h1 class="text-accent text-center text-5xl font-bold">{$page.status}</h1> <!-- http error code --></div>
<!-- redis llvmcimpoiel to arm asm archiecut rhetpt status codes tcp realtime stream to stream data to a device from another in real time -->
<div class="divider"><p class = "text-accent text-center text-4xl mt-10">{$page.error.message}</p></div> <!-- http error message -->
<br>
<p class = "text-accent text-center text-4xl font-bold mt-10">{jwt}</p>
<br>

<p class="text-accent text-center text-xl mt-10"><a href="/">Go Home</a></p>
</body>