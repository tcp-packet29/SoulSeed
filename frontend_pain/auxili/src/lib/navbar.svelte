<script>
    export let id = ""
</script>

<div class="navbar bg-base-100 bg-accent z-50 top-0 w-full absolute sticky">
    <div class="flex-1">
        <h1 class="font-bold text-2xl text-accent m-5">🌱 Soulseed</h1>
    </div>
    <div class="">
        <div class="join">
            <button class="join-item btn btn-secondary text-accent"><a class="link-accent" href="/login">Login</a></button>
            <button class="join-item btn btn-secondary text-accent"><a class="link-accent" href="/register">Register</a></button>
            <button class="join-item btn btn-secondary text-accent"><a class="link-accent" href="/organization">Organizations</a></button>
            <button class="join-item btn btn-secondary text-accent"><a class="link-accent" href="/createOrg">Create an Organization</a></button>
            {#if id != ""}
<!--                <button class="join-item btn btn-secondary text-accent"><div class="divider"></div></button>-->
                <button class="join-item btn btn-secondary text-accent" ><a class="link-accent" href="/mainPg/{id}">Check Trades</a></button>
            {/if}
        </div>
    </div>
    <!--    <div class="flex-none">-->

    <!--            <img class="mask mask-circle" src="data:image/jpeg;base64,/9j/4AAQSkZJRgABAQAAAQABAAD/2wCEAAkGBwgHBgkIBwgKCgkLDRYPDQwMDRsUFRAWIB0iIiAdHx8kKDQsJCYxJx8fLT0tMTU3Ojo6Iys/RD84QzQ5OjcBCgoKDQwNGg8PGjclHyU3Nzc3Nzc3Nzc3Nzc3Nzc3Nzc3Nzc3Nzc3Nzc3Nzc3Nzc3Nzc3Nzc3Nzc3Nzc3Nzc3N//AABEIAIAAgAMBIgACEQEDEQH/xAAaAAEAAgMBAAAAAAAAAAAAAAAABgcDBAUC/8QAMhAAAgIBAgIIAwgDAAAAAAAAAAECAwQFEQYxEiFRYXGRocETYtEiIzNBQlNjsTJykv/EABYBAQEBAAAAAAAAAAAAAAAAAAABAv/EABYRAQEBAAAAAAAAAAAAAAAAAAARAf/aAAwDAQACEQMRAD8AtIAGmQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMOZlV4WLbk3PauqPSff3AYdS1LF02n4mVZtv/jCPXKXgiL5XGWVKTWLjVVx/k3k35bI4OoZt2oZU8jIf2pcorlFdiNYsKk+LxllRkllY1Vkf494tee5KNN1LF1Kn4mLZvt/lCXVKHiisDZ0/Nu0/KhkY72lHmnykuxgWkDDh5NeZi1ZNL3rtj0kZiAAAAAAAAAAABGeOsiUMHHoi9vi2Nvwivq0SYinHsfusKf5KU16L6AQ8AGkAATVxNeBciU8HIok9/h2Jrwa+qZJiKcBR+6zZ/k5QX9/UlZAAAAAAAAAAAA5HFOFLN0ixVrpWVNWRS5vbn6bnXAFSJ78gS/XOFpTsnk6Z0ftPeVDe3X8r9iNX6dnUSauw747fI9vMtI1RvtzNqnTs6+SVWHfJv5Ht58iS6HwtKFkMnU+j9l7xoT36/mfsB1eFsKWFpFasXRstfxJJ81vyXlsdcAgAAAAAAAAAAADHfdXj0yuvnGFcFvKUuSIfq3Fl1snXpsfhV/uyW8peHZ/YEwvvpx49O+2uqPbOSivU58+IdKrbTzYP/RN+qRXd1tl83O6ydk3+qcnJ+p4KVZFfEOlWNbZsE/nTX9o36L6ciPTotrtj2wkpL0KoPdNtlE1Omydc1+qEnF+gSrYBC9J4surlGvUo/Fr/AHYraUfHtJhRdXkUwuonGdc1vGUeTIrIAAAAAAAAeLbIU1yttkoVwW8pPkkeyJcb6i10NOqe26U7e/sXv5AcXXtZs1bI6t440H93X7vvOWAU0ABUAAAOpoOs2aTkde8sab+8r913nLBFWxVZC6uNtUlOua3jJcmj2RLgjUW+np1r32TnV7r38yWkAAAAAA8eRVupZMszPyMhv8Sxtdy36l5bFk6lN1adl2J7OFM2v+WVauRQABUAAAAAAAAbOm5Lw8/HyE/w7E33r815blorrXUVK+RaWm2O3TsWxvdzpg349FEVsgAg/9k=">-->

    <!--    </div>-->
    <!--    <a class="btn btn-ghost btn-sm rounded-btn bg-secondary" href="http://localhost:5173">Auxili</a>-->
    <!--    <h1>&#45;&#45;</h1>-->
    <!--    <h2><a class="btn btn-sm btn-ghost rounded-btn bg-secondary" href="http://localhost:5173/register">Register</a></h2>-->
    <!--    <h1>&#45;&#45;</h1>-->
    <!--    <h2><a class="btn btn-ghost btn-sm rounded-btn bg-secondary" href="http://localhost:5173/login">Login</a></h2>-->
    <!--    &lt;!&ndash; llvmjwt &ndash;&gt;-->
</div>