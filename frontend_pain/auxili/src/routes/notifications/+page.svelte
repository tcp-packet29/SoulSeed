<script>
    import {getToken} from '$lib/parse.svelte'
    import axios from 'axios'
    import {browser} from '$app/environment'


    axios.get('http://localhost:8080/app/notifs', {
        headers: {
            "Token": getToken(),
            "Content-Type": "application/json",
        }}
    ).then((response) => {
        console.log(response)

    }).catch(function(error) {
        //not the    erorr funciton but still okay arrow
        console.log(error)
    })

    axios.post('http://localhost:8080/app/notifs/add', {
        "UserFor": "64b4bbe5772745e4fd8775a3",
        "Body": "LLVM IR JWT AUTHENTICATION MIDDLEWARE EXTENDING LIP THROUGH MACROS ",
    }, {
        headers: {
            "Token": getToken(),
            "Content-Type": "application/json",
        },

    }).then((response)=> {
        console.log(response)
    }).catch(f  unction(err) {
        console.log(err)
    })


    // <!--&lt;!&ndash;&lt;!&ndash;axios.get('http://localhost:8080/access/users/token', {&ndash;&gt;&ndash;&gt;-->
    // // <!--&lt;!&ndash;    headers: {&ndash;&gt;-->
    // // // <!--        "Token": getToken(),-->
    // // // <!--        "Content-Type": "application/json",-->
    // // // <!--    }}-->
    // // // <!--).then((response) => {-->
    // // // <!--    console.log(response)-->
    // // // <!--}).catch(function(error) {-->
    // // // <!--    //not the erorr funciton but still okay arrow-->
    // // // <!--    console.log(error)-->
    // // // // })
    // // //
    //macro preprocessor

</script>

<!-- new html dialog tag for nativebrowser rendering -->