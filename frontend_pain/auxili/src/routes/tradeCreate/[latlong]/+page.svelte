<script>
    import { page } from '$app/stores'
    import axios from 'axios';
    let x = $page.params.latlong.split('~')
    console.log(x)

    let lat = x[0]
    let liblong = x[2   ]

    axios.get(' https://api.mapbox.com/geocoding/v5/mapbox.places/' + liblong + ',' + lat + '.json?types=postcode&limit=1&access_token=pk.eyJ1Ijoic3BlbGxjYXN0IiwiYSI6ImNsZTN1YjNtcTBjaGczb2xmMzJ1YnZua2IifQ.ZzbJhqpfTl94WAt8jpUHvA')
        .then(function (response) {
            console.log(response.data.features[0].text);

        })
        .catch(function (error) {
            console.log(error);
        })

    


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
    <div data-theme="mycodecompiled" class="card w-96 bg-secondary bg-base-200 text-neutral-content shadow-xl">
        <div class="card-body items-center text-center">
            <h2 class="card-title text-neutral tooltip" data-tip={lat + ", " + liblong}>Create Trade</h2>

        </div>
    </div>
</div>