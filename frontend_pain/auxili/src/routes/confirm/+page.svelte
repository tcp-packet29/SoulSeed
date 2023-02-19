<script>

    import { browser } from '$app/environment';
    if (browser) {
        try {
            let token = window.localStorage.getItem("conf")
        } catch {
            console.log('couldnt fetch')
        }
    }



</script>