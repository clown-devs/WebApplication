window.onload = async function () {
    let a = document.getElementById('bronir');
    let b = document.getElementById('savethisshit1')
    let c = document.getElementById('bronir3')
    let r = document.getElementById('add')
    let q = document.getElementById('contacts')
    let close = document.getElementById('close')
    let client = document.getElementById('cliente')
    let place = document.getElementById('place_list')
    let responce = await fetch('http://176.119.157.82:8000/api/meeting/')
    let content = await responce.json()
    if (content == 0) {
        let list = document.querySelector('.current_nav_main')

        list.innerHTML = `
            <a href="#zatemnenie_edit_klient"><img class="pencil" src="./img/pencil.png" alt=""></a>
           <div class="current_nav_elem_themee ">ВСТРЕЧ НЕТ</div>

`
        let listt = document.querySelector('.meet_list_list')
        listt.innerHTML = `
        <div class="meet_list_elem_dataa">Встреч нет</div>
        `

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
            let str1 = content[key].end
            str = str.substr(0, 5)
            str1 = str1.substr(0, 5);

            list.innerHTML += `
                
           <li class="meet_list_elem">
           
                    <div class="meet_list_elem_data">${content[key].date}</div>
                    <div class="meet_list_elem_time">${str}</div>
                    <div class="meet_list_elem_time1">${str1}</div>
                    <div class="meet_list_elem_place">${content2.name}</div>
                    <div class="meet_list_elem_theme">${content[key].topic}</div> 
                    <div class="meet_list_elem_klient">${content3.company}</div>
           </li>   
           `
        }

        document.getElementById('btn_current').onclick = function() {
            list.innerHTML = '';

            for (key in content) {
                let str = content[key].start
                str = str.substr(0, 5)
                let str1 = content[key].end
                str1 = str1.substr(0, 5);

                list.innerHTML += `
                
           <li class="meet_list_elem">
           
                    <div class="meet_list_elem_data">${content[key].date}</div>
                    <div class="meet_list_elem_time">${str}</div>
                    <div class="meet_list_elem_time1">${str1}</div>
                    <div class="meet_list_elem_place">${content2.name}</div>
                    <div class="meet_list_elem_theme">${content[key].topic}</div> 
                    <div class="meet_list_elem_klient">${content3.company}</div>
           </li>   
           `
            }
        };

        document.getElementById('btn_past').onclick = function() {

            list.innerHTML = `
                
           <li class="meet_list_elem">
           
                    <div class="meet_list_elem_data">2019-05-21</div>
                    <div class="meet_list_elem_time">15:30</div>
                    <div class="meet_list_elem_time1">17:30</div>
                    <div class="meet_list_elem_place">Переговорка 1.2</div>
                    <div class="meet_list_elem_theme">Тест тема</div> 
                    <div class="meet_list_elem_klient">Yandex</div>
           </li>   
           `

        };

        
        let list_blizh = document.querySelector('.current_nav_main')
        let str = content[0].start
        let str1 = content[0].end
        str = str.substr(0, 5)
        str1 = str1.substr(0, 5);
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
            <div class="current_nav_elem_times1">${str1}</div>
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
<<<<<<< HEAD

=======
        console.log(data);
        console.log(theme);
>>>>>>> c5fa132570fbfaf141118d5a4f2e7f3824ddcfb1
        let body = {

            date: data.value,
            start: time_start.value,
            end: time_end.value,
            topic: theme.value,
            creator: '1',
            place: id_place,
            client: id_client,
            contact: id_contact
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

        // window.location = "http://sbermeeting.tk/"
    }

    
    let k = 0;
    let i = 0;
    let mark = 1;
    r.onclick = async function () {

        let client = await fetch('http://176.119.157.82:8000/api/client/')
        let contacts = await fetch('http://176.119.157.82:8000/api/contact/')
        let places = await fetch('http://176.119.157.82:8000/api/place/')
        let content1 = await client.json()
        let content2 = await contacts.json()
        let content3 = await places.json()
        let list_dolj = document.querySelector('.dolj_input1')
        let list_cli = document.querySelector('.klient_input')
        let list_contact = document.querySelector('.litso_input')
        let list_place = document.querySelector('.place_input')

        let keyS;
        
        if (mark == 1) {
            for (keyS in content1) {
                list_cli.innerHTML += `
            
            <option value="${content1[keyS].company}" id="company">${content1[keyS].company}</option>
            
           `
            }
            for (keyS in content2) {
                list_contact.innerHTML += `
            
            <option value="${content2[keyS].first_name + ' ' + content2[keyS].second_name + ' ' + content2[keyS].third_name}" id="nam">${content2[keyS].first_name + ' ' + content2[keyS].second_name + ' ' + content2[keyS].third_name}</option>
           `

            }
            list_dolj.innerHTML += `
           <input class="dolj_input" type="text" id="dolj" value="${content2[0].position}">
               `
               for (keyS in content3) {
                list_place.innerHTML += `
                
                <option value="${content3[keyS].name}" id="qwe">${content3[keyS].name}</option>           `
            }
        }


        mark = 0;
    }


    // let select = document.querySelector('.klient_input')
    // let op = document.querySelector('option')
    // op.addEventListener('click', function() {
    //     console.log(selector.value);
    // })
    
    let id_contact = 1
    let id_client = 1
    let id_place = 1
    q.onclick = async function () {
        
        let list_dolj = document.querySelector('.dolj_input1')
        let contacts = await fetch('http://176.119.157.82:8000/api/contact/')
        let content2 = await contacts.json()
        let name = document.getElementById('contacts')

        let name1 = (content2[0].first_name + ' ' + content2[0].second_name + ' ' + content2[0].third_name)

        if ((name.value) == name1) {
            id_contact = 1
            list_dolj.innerHTML = `
                   <input class="dolj_input" type="text" id="dolj" value="${content2[0].position}">
                       `
        }
        else {
            id_contact = 2
            list_dolj.innerHTML = `
                   <input class="dolj_input" type="text" id="dolj" value="${content2[1].position}">
                       `
        }
<<<<<<< HEAD
    }

    client.onclick = async function () {
        let client = await fetch('http://176.119.157.82:8000/api/client/')
        let content1 = await client.json()

        let name = document.getElementById('cliente')

        let name1 = (content1[0].company)

        if ((name.value) == name1) {
            id_client = 1
            
        }
        else {
            id_client = 2
            
        }
    }

    place.onclick = async function () {
        let places = await fetch('http://176.119.157.82:8000/api/place/')
        let content3 = await places.json()

        let name = document.getElementById('place_list')

        let name1 = (content3[0].name)

        if ((name.value) == name1) {
            id_place = 1
            
        }
        else {
            id_place = 2
            
        }
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

=======
>>>>>>> c5fa132570fbfaf141118d5a4f2e7f3824ddcfb1
    }
    getResponse()
}

    client.onclick = async function () {
        let client = await fetch('http://176.119.157.82:8000/api/client/')
        let content1 = await client.json()

        let name = document.getElementById('cliente')

        let name1 = (content1[0].company)

        if ((name.value) == name1) {
            id_client = 1
            
        }
        else {
            id_client = 2
            
        }
    }

    place.onclick = async function () {
        let places = await fetch('http://176.119.157.82:8000/api/place/')
        let content3 = await places.json()

        let name = document.getElementById('place_list')

        let name1 = (content3[0].name)

        if ((name.value) == name1) {
            id_place = 1
            
        }
        else {
            id_place = 2
            
        }
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




}