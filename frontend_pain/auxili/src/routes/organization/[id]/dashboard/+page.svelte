<script>
    import { page } from '$app/stores'
    import axios from 'axios'
    import { browser } from '$app/environment'
    import Chart from 'chart.js/auto'
    import { getToken } from '$lib/parse.svelte'

    let data;




    if (browser) {
        axios.get('http://localhost:8080/access/allstuff')
            .then(function (res) {
                console.log(res.data.data.data)
                data = {
                    labels: Object.keys(res.data.data.data),
                    datasets: [
                        {
                            label: 'Variety of jwt auth',
                            data: Object.values(res.data.data.data)
                        }
                    ],
                    borderWidth: 9
                }
                new Chart(document.getElementById("jwt"), {
                    type: 'bar',
                    data: data,
                    options: {
                        scales: {
                            y: {
                                beginAtZero: true
                            }
                        }
                    }
                })
            })
            .catch(function (err) {

            })
    }

</script>

<button class="btn btn-primary">
    <a href="/organization/{$page.params.id}/invite/{$page.params.id}">Home</a>

</button>

<div class="w-1/2 p-5 rounded-lg justify-center content-center text-center udp amqp gorm chart.js">
    <canvas id="jwt"></canvas>
</div>