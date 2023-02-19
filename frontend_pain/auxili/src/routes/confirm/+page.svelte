<script>

    import { browser } from '$app/environment';

    let val = ""
    let token = ""
    if (browser) {
        try {
            token = window.localStorage.getItem("conf")
            console.log(token)
        } catch {
            console.log('couldnt fetch')
        }
    }

    function Check() {
        if (val === token) {
            console.log('correct')
        } else {
            console.log('incorrect')
        }
    }
//  could just look in local storage
    //shpuld proabably add usr here but okay



</script>

<div class="flex h-screen justify-center items-center">
    <div data-theme="mycodecompiled" class="card w-96 bg-secondary bg-base-200 text-neutral-content shadow-xl">
        <div class="card-body items-center text-center">
            <h2 class="card-title text-neutral">Confirmation</h2>
            <p class="mb-5">Please enter your confirmation code that was sent to your email.</p>
            <input type="password" placeholder="Password" class="input input-bordered input-accent w-full max-w-xs" bind:value={val}/>
            <div class="card-actions">
                <button class="btn btn-outline btn-accent" on:click={Check}>Check</button>
            </div>
        </div>
    </div>
</div>