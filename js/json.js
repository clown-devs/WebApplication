window.onload = async function () {
    let a = document.getElementById('bronir');
    let b = document.getElementById('savethisshit')
    let c = document.getElementById('bronir3')
    let r = document.getElementById('cliente')

    let responce = await fetch('http://176.119.157.82:8000/api/meeting/')
    let content = await responce.json()
    if (content == 0) {
        let list = document.querySelector('.current_nav_main')

        list.innerHTML = `
            <a href="#zatemnenie_edit_klient"><img class="pencil" src="./img/pencil.png" alt=""></a>
           <div class="current_nav_elem_themee ">ВСТРЕЧ НЕТ</div>

`
        // let listt = document.querySelector('.meet_list_list')
        // listt.innerHTML = `
        // <div class="meet_list_elem_dataa">Встреч нет</div>
        // `

    }
    else {
        let id_contact = content[0].contact
        let id_place = content[0].place
        let id_client = content[0].client
        let responce1 = await fetch('http://176.119.157.82:8000/api/contact/' + id_contact)
        let responce2 = await fetch('http://176.119.157.82:8000/api/place/' + id_place)
        let responce3 = await fetch('http://176.119.157.82:8000/api/client/' + id_client)
        let content1 = await responce1.json()
        let content2 = await responce2.json()
        let content3 = await responce3.json()

        let list = document.querySelector('.meet_list_list')

        let key;

        for (key in content) {
            let str = content[key].start
            str = str.substr(0, 5)

            list.innerHTML += `
                
           <li class="meet_list_elem">
           
                    <div class="meet_list_elem_data">${content[key].date}</div>
                    <div class="meet_list_elem_time">${str}</div>
                    <div class="meet_list_elem_place">${content2.name}</div>
                    <div class="meet_list_elem_theme">${content[key].topic}</div> 
                    <div class="meet_list_elem_klient">${content3.company}</div>
           </li>   
           `
        }

        let list_blizh = document.querySelector('.current_nav_main')
        let str = content[0].start
        str = str.substr(0, 5)
        list_blizh.innerHTML += `
            
            <div class="current_nav_elem_klient">${content3.company}</div>
            <div class="current_nav_elem_Kont_litso">${content1.second_name + " " + content1.first_name + " " + content1.third_name}</div>
            <div class="current_nav_elem_phone">${content1.phone}</div>
            <div class="current_nav_elem_theme">Тема встречи:</div>
            <div class="current_nav_elem_themes">${content[0].topic}</div>
            <div class="current_nav_elem_data">Дата:</div>
            <div class="current_nav_elem_datas">${content[0].date}</div>
            <div class="current_nav_elem_place">Место:</div>
            <div class="current_nav_elem_places">${content2.name}</div>
            <div class="current_nav_elem_time">Время:</div>
            <div class="current_nav_elem_times">${str}</div>
            `
    }

    b.onclick = async function () {
        async function sendRequest(method, url, body = null) {
            return fetch(url, {
                method: method,
                body: JSON.stringify(body),
                headers: {
                    'Content-type': 'application/json; charset=UTF-8',
                }

            })

        }
        let data = document.getElementById("data")
        let time_start = document.getElementById("time_start")
        let time_end = document.getElementById("time_end")
        let theme = document.getElementById("theme")
        let place = document.getElementById("place")
        let client = document.getElementById("client")
        let contact_lico = document.getElementById("contact_lico")
        console.log(data);
        let body = {

            date: data.value,
            start: time_start.value,
            end: time_end.value,
            topic: theme.value,
            creator: '1',
            place: place.value,
            client: client.value,
            contact: contact_lico.value
            // date: '2022-11-14',
            // start: '11:34',
            // end: '18:45',
            // topic: 'Проблемы будущих нас',
            // creator: '2',
            // place: '2',
            // client: '2',
            // contact: '2'

        }

        let responce = await sendRequest('POST', 'http://176.119.157.82:8000/api/meeting/', body)
        responce = await responce.json()


        let id_clien = responce.creator
        let id_meet = responce.id
        console.log(responce)
        console.log(id_meet)
        fetch('http://176.119.157.82:8000/api/employeelist/', {
            method: 'POST',
            body: JSON.stringify({
                "employee": id_clien,
                "meeting": id_meet
            }),
            headers: {
                'Content-type': 'application/json; charset=UTF-8',
            },
        })
            .then((response) => response.json())
            .then((json) => console.log(json));

        window.location = "http://sbermeeting.tk/";
    }




    r.onclick = async function () {
    let responcee = await fetch('http://176.119.157.82:8000/api/client/')
    let contente = await responcee.json()    
    let list_cli = document.querySelector('.klient_input')
    let keyS;

        for (keyS in contente) {
            list_cli.innerHTML += `
            
            <option value="one" id="valu">${contente[keyS].company}</option>
           `
        }



    // let select = document.querySelector('.klient_input')
    // let op = document.querySelector('option')
    // op.addEventListener('click', function() {
    //     console.log(selector.value);
    // })
        }






    // a.onclick = async function () {

    //     async function sendRequest(method, url, body = null) {
    //         return fetch(url, {
    //             method: method,
    //             body: JSON.stringify(body),
    //             headers: {
    //                 'Content-Type': 'application/json'
    //             }

    //         })

    //     }
    //     let body = {
    //         color: 'YES'

    //     }

    //     let responce = await sendRequest('PATCH', 'http://176.119.157.82:8000/api/v1/cars/car/detail/9/', body)

    //     let content = await responce.json()
    //     console.log(content)


    //     let list = document.querySelector('.posts')

    //     let key
    //     list.innerHTML = `
    //         <li class="post">
    //         <h4>${content.id}</h4>
    //         </li>   
    //         <li class="post"> 
    //             <h4>${content.color}</h4>
    //         </li>`

    //     let if_id = "list"

    //     if_id = document.getElementById(if_id);
    //     if_id.style.background = "rgb(218, 64, 92)";



    // }


    //  c.onclick = async function () {


    {/* <img class="pencil" src="pencil.png" alt=""> */ }
    // let if_id = "list"
    // if (content.color == "YES") {
    //     if_id = document.getElementById(if_id);
    //     if_id.style.background = "rgb(218, 64, 92)";
    // }
    // else {
    //     if_id = document.getElementById(if_id);
    //     if_id.style.background = "rgb(137, 206, 126)";
    // }
    //  }



    async function update() {
        let responce = await fetch('http://176.119.157.82:8000/api/v1/cars/car/detail/9/')

        let content = await responce.json()

        let list = document.querySelector('.posts')

        let key;

        for (key in content) {
            console.log(content[key])
            list.innerHTML = `
   <li class="post">
   <h4>${content.id}</h4>
   </li>   
   <li class="post"> 
       <h4>${content.color}</h4>
   </li>`

        }
        getResponse()
    }

}