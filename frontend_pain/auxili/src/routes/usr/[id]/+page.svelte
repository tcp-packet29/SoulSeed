<script>
    import { page } from '$app/stores';
    import axios from 'axios';



    console.log($page.params.id)


    axios.get("http://localhost:8080/access/users/" + $page.params.id).then((response) => {
        if (response.status == 200) {
            let usr = JSON.parse(JSON.stringify(response.data))
            console.log(usr, usr.data.data.Username)

        }

    }).catch((error) => {
        console.log(error)

    })





</script>
