// window.onload = async function () {
//     let a = document.getElementById('bronir');
//     let c = document.getElementById('bronir3')
//     let r = document.getElementById('add')
    
//     let save_button = document.getElementById('savethisshit1')
//     let contact_button = document.getElementById('contacts')
//     let client_button = document.getElementById('cliente')
    
//     let close = document.getElementById('close')
//     let place = document.getElementById('place_list')
    
//     let responce = await fetch('http://176.119.157.82:8000/api/meeting/')
//     let content = await responce.json()
    
//     if (content == 0) {
//         let list = document.querySelector('.current_nav_main')

//         list.innerHTML = `
//             <a href="#zatemnenie_edit_klient"><img class="pencil" src="./img/pencil.png" alt=""></a>
//            <div class="current_nav_elem_themee ">ВСТРЕЧ НЕТ</div>
// `
//         let listt = document.querySelector('.meet_list_list')
//         listt.innerHTML = `
//         <div class="meet_list_elem_dataa">Встреч нет</div>
//         `
//     }

//     else {
//         let id_contact = content[0].contact
//         let id_place = content[0].place
//         let id_client = content[0].client
        
//         let responce1 = await fetch('http://176.119.157.82:8000/api/contact/' + id_contact)
//         let responce2 = await fetch('http://176.119.157.82:8000/api/place/' + id_place)
//         let responce3 = await fetch('http://176.119.157.82:8000/api/client/' + id_client)
        
//         let content1 = await responce1.json()
//         let content2 = await responce2.json()
//         let content3 = await responce3.json()

//         let list = document.querySelector('.meet_list_list')

//         for (key in content) {
//             if (key == 0) {
//                 continue
//             }
            
//             let str = content[key].start
//             let str1 = content[key].end
//             str = str.substr(0, 5)
//             str1 = str1.substr(0, 5);
        
//             list.innerHTML += `
                    
//             <li class="meet_list_elem">
            
//                         <div class="meet_list_elem_data">${content[key].date}</div>
//                         <div class="meet_list_elem_time">${str}</div>
//                         <div class="meet_list_elem_time1">${str1}</div>
//                         <div class="meet_list_elem_place">${content[key].place}</div>
//                         <div class="meet_list_elem_theme">${content[key].topic}</div> 
//                         <div class="meet_list_elem_klient">${content[key].client}</div>
//             </li>   
//             `
//         }

//         document.getElementById('btn_current').onclick = function() {
//         //     list.innerHTML = '';

//         //     for (key in content) {
//         //         let str = content[key].start
//         //         str = str.substr(0, 5)
//         //         let str1 = content[key].end
//         //         str1 = str1.substr(0, 5);

//         //         list.innerHTML += `
                
//         //    <li class="meet_list_elem">
           
//         //             <div class="meet_list_elem_data">${content[key].date}</div>
//         //             <div class="meet_list_elem_time">${str}</div>
//         //             <div class="meet_list_elem_time1">${str1}</div>
//         //             <div class="meet_list_elem_place">${content2.name}</div>
//         //             <div class="meet_list_elem_theme">${content[key].topic}</div> 
//         //             <div class="meet_list_elem_klient">${content3.company}</div>
//         //    </li>   
//         //    `
//         //     }
//         };

//         document.getElementById('btn_past').onclick = function() {

//             list.innerHTML = `
                
//            <li class="meet_list_elem">
           
//                     <div class="meet_list_elem_data">2019-05-21</div>
//                     <div class="meet_list_elem_time">15:30</div>
//                     <div class="meet_list_elem_time1">17:30</div>
//                     <div class="meet_list_elem_place">Переговорка 1.2</div>
//                     <div class="meet_list_elem_theme">Тест тема</div> 
//                     <div class="meet_list_elem_klient">Yandex</div>
//            </li>   
//            `

//         };

        
//         let list_blizh = document.querySelector('.current_nav_main')
//         let str = content[0].start
//         let str1 = content[0].end
//         str = str.substr(0, 5)
//         str1 = str1.substr(0, 5);
//         list_blizh.innerHTML += `
            
//             <div class="current_nav_elem_klient">${content3.company}</div>
//             <div class="current_nav_elem_Kont_litso">${content1.second_name + " " + content1.first_name + " " + content1.third_name}</div>
//             <div class="current_nav_elem_phone">${content1.phone}</div>
//             <div class="current_nav_elem_theme">Тема встречи:</div>
//             <div class="current_nav_elem_themes">${content[0].topic}</div>
//             <div class="current_nav_elem_data">Дата:</div>
//             <div class="current_nav_elem_datas">${content[0].date}</div>
//             <div class="current_nav_elem_place">Место:</div>
//             <div class="current_nav_elem_places">${content2.name}</div>
//             <div class="current_nav_elem_time">Время:</div>
//             <div class="current_nav_elem_times">${str}</div>
//             <div class="current_nav_elem_times1">${str1}</div>
//             `
//     }
    
//     save_button.onclick = async function() {
//         async function sendRequest(method, url, body = null) {
//             return fetch(url, {
//                 method: method,
//                 body: JSON.stringify(body),
//                 headers: {
//                     'Content-type': 'application/json; charset=UTF-8',
//                 }

//             })

//         }

//         let time_start = document.getElementById("time_start")
//         let time_end = document.getElementById("time_end")
//         let theme = document.getElementById("theme")
//         let data = document.getElementById("data")

//         let body = {

//             date: data.value,
//             start: time_start.value,
//             end: time_end.value,
//             topic: theme.value,
//             creator: '1',
//             place: id_place,
//             client: id_client,
//             contact: id_contact,
//             creatorText: ''
//             // date: '2022-11-14',
//             // start: '11:34',
//             // end: '18:45',
//             // topic: 'Проблемы будущих нас',
//             // creator: '2',
//             // place: '2',
//             // client: '2',
//             // contact: '2'

//         }

//         let responce = await sendRequest('POST', 'http://176.119.157.82:8000/api/meeting/', body)
//         responce = await responce.json()


//         let id_clien = responce.creator
//         let id_meet = responce.id

//         fetch('http://176.119.157.82:8000/api/employeelist/', {
//             method: 'POST',
//             body: JSON.stringify({
//                 "employee": id_clien,
//                 "meeting": id_meet
//             }),
//             headers: {
//                 'Content-type': 'application/json; charset=UTF-8',
//             },
//         })

//         //  window.location = "http://sbermeeting.tk/"
//     }

//     let id_client = await get_first_id("client")
//     client_button.onclick = async function() {
//         this.selected_client
        
//         let client = document.getElementById('cliente').addEventListener(
//             'change',
//             () => { this.selected_client = this.value},
//             {'once' : true}
//         )

//         id_client = await get_client_id(this.selected_client)
//     }

//     let id_contact = await get_first_id("contact")
//     contact_button.onclick = async function() {
//         this.selected_contact
        
//         let contact = document.getElementById('contacts').addEventListener(
//             'change',
//             () => { this.selected_contact = this.value},
//             {'once' : true}
//         )
//         id_contact = await get_contact_id(this.selected_contact)
//     }

//     let id_place = await get_first_id("place")
//     place.onclick = async function () {
//         this.selected_place
        
//         let place = document.getElementById('place_list').addEventListener(
//             'change',
//             () => { this.selected_place = this.value},
//             {'once' : true}
//         )
//         id_place = await get_place_id(this.selected_place)
//     }

//     let mark = 1;
//     r.onclick = async function () {

//         let client = await fetch('http://176.119.157.82:8000/api/client/')
//         let contacts = await fetch('http://176.119.157.82:8000/api/contact/')
//         let places = await fetch('http://176.119.157.82:8000/api/place/')
        
//         let content1 = await client.json()
//         let content2 = await contacts.json()
//         let content3 = await places.json()
        
//         let list_dolj = document.querySelector('.dolj_input1')
//         let list_cli = document.querySelector('.klient_input')
//         let list_contact = document.querySelector('.litso_input')
//         let list_place = document.querySelector('.place_input')

//         let keyS;
        
//         if (mark == 1) {
//             for (keyS in content1) {
//                 list_cli.innerHTML += `
            
//             <option value="${content1[keyS].company}" id="company">${content1[keyS].company}</option>
            
//            `
//             }
//             for (keyS in content2) {
//                 list_contact.innerHTML += `
            
//             <option value="${content2[keyS].first_name + ' ' + content2[keyS].second_name + ' ' + content2[keyS].third_name}" id="nam">${content2[keyS].first_name + ' ' + content2[keyS].second_name + ' ' + content2[keyS].third_name}</option>
//            `

//             }
//             list_dolj.innerHTML += `
//            <input class="dolj_input" type="text" id="dolj" value="${content2[0].position}">
//                `
//                for (keyS in content3) {
//                 list_place.innerHTML += `
                
//                 <option value="${content3[keyS].name}" id="qwe">${content3[keyS].name}</option>           `
//             }
//         }


//         mark = 0;
//     }

// }

// async function get_contact_id(contact_full_name) {

//     let responce_contacts = await fetch('http://176.119.157.82:8000/api/contact/')
//     let contacts = await responce_contacts.json()
//     let id_contact = get_first_id("contact")

//     contacts.forEach((contact) => {
//         let full_name = contact.first_name + ' '
//             + contact.second_name + ' '
//             + contact.third_name

//         if (full_name === contact_full_name) {
//             id_contact = contact.id
//         }
//     });

//     return id_contact
// }

// async function get_client_id(client_company) {

//     let responce = await fetch('http://176.119.157.82:8000/api/client/')
//     let clients = await responce.json()
//     let id_client = get_first_id("client")

//     clients.forEach((client) => {
//         let company_name = client.company

//         if (client_company === company_name) {
//             id_client = client.id
//         }
//     });

//     return id_client
// }

// async function get_place_id(selected_place) {
    
//     let responce = await fetch('http://176.119.157.82:8000/api/place/')
//     let places = await responce.json()
//     let id_place = get_first_id("place")

//     places.forEach((place) => {
//         let place_name = place.name

//         if (place_name === selected_place) {
//             id_place = place.id
//         }
//     });

//     return id_place 
// }

// async function get_first_id(api_path) {
//     let responce = await fetch('http://176.119.157.82:8000/api/' + api_path)
//     let objects = await responce.json()
//     return objects[0].id
// }

window.onload = async function () {
    let a = document.getElementById('bronir');
    let c = document.getElementById('bronir3')
    let r = document.getElementById('add')
    
    let save_button = document.getElementById('savethisshit1')
    let contact_button = document.getElementById('contacts')
    let client_button = document.getElementById('cliente')
    
    let close = document.getElementById('close')
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

        display_meetings("future")

        document.getElementById('btn_current').onclick = function() {
            display_meetings("future")
        };

        document.getElementById('btn_past').onclick = function() {
            display_meetings("past")
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
    
    save_button.onclick = async function() {
        async function sendRequest(method, url, body = null) {
            return fetch(url, {
                method: method,
                body: JSON.stringify(body),
                headers: {
                    'Content-type': 'application/json; charset=UTF-8',
                }

            })

        }

        let time_start = document.getElementById("time_start")
        let time_end = document.getElementById("time_end")
        let theme = document.getElementById("theme")
        let data = document.getElementById("data")

        let body = {

            date: data.value,
            start: time_start.value,
            end: time_end.value,
            topic: theme.value,
            creator: '1',
            place: id_place,
            client: id_client,
            contact: id_contact,
            creatorText: ''
            // date: '2022-11-14',
            // start: '11:34',
            // end: '18:45',
            // topic: 'Проблемы будущих нас',
            // creator: '2',
            // place: '2',
            // client: '2',
            // contact: '2'
            //
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

        //  window.location = "http://sbermeeting.tk/"
    }

    let id_client = await get_first_id("client")
    client_button.onclick = async function() {
        this.selected_client
        
        let client = document.getElementById('cliente').addEventListener(
            'change',
            () => { this.selected_client = this.value},
            {'once' : true}
        )

        id_client = await get_client_id(this.selected_client)
    }

    let id_contact = await get_first_id("contact")
    contact_button.onclick = async function() {
        this.selected_contact
        
        let contact = document.getElementById('contacts').addEventListener(
            'change',
            () => { this.selected_contact = this.value},
            {'once' : true}
        )
        id_contact = await get_contact_id(this.selected_contact)
    }

    let id_place = await get_first_id("place")
    place.onclick = async function () {
        this.selected_place
        
        let place = document.getElementById('place_list').addEventListener(
            'change',
            () => { this.selected_place = this.value},
            {'once' : true}
        )
        id_place = await get_place_id(this.selected_place)
    }

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

}

async function get_contact_id(contact_full_name) {

    let responce_contacts = await fetch('http://176.119.157.82:8000/api/contact/')
    let contacts = await responce_contacts.json()
    let id_contact = get_first_id("contact")

    contacts.forEach((contact) => {
        let full_name = contact.first_name + ' '
            + contact.second_name + ' '
            + contact.third_name

        if (full_name === contact_full_name) {
            id_contact = contact.id
        }
    });

    return id_contact
}

async function get_client_id(client_company) {

    let responce = await fetch('http://176.119.157.82:8000/api/client/')
    let clients = await responce.json()
    let id_client = get_first_id("client")

    clients.forEach((client) => {
        let company_name = client.company

        if (client_company === company_name) {
            id_client = client.id
        }
    });

    return id_client
}

async function get_place_id(selected_place) {
    
    let responce = await fetch('http://176.119.157.82:8000/api/place/')
    let places = await responce.json()
    let id_place = get_first_id("place")

    places.forEach((place) => {
        let place_name = place.name

        if (place_name === selected_place) {
            id_place = place.id
        }
    });

    return id_place 
}

async function get_first_id(api_path) {
    let responce = await fetch('http://176.119.157.82:8000/api/' + api_path)
    let objects = await responce.json()
    return objects[0].id
}

async function display_meetings(type) {
    
    let path = ""
    if (type == "past") {
        path = "?past=true"
    }

    let responce = await fetch('http://176.119.157.82:8000/api/meeting' + path)
    let meetings = await responce.json()
    
    let list = document.querySelector('.meet_list_list')
    
    responce = await fetch('http://176.119.157.82:8000/api/place/')
    let places = await responce.json()

    responce = await fetch('http://176.119.157.82:8000/api/client/')
    let clients = await responce.json()

    list.innerHTML = '';
    for (key in meetings) {
        if (key == 0) { continue; }

        let str = meetings[key].start
        str = str.substr(0, 5)
        
        let str1 = meetings[key].end
        str1 = str1.substr(0, 5);
        
        let place
        places.forEach((elem) => {
            if (elem.id === meetings[key].place) {
                place = elem.name
            }
        });

        let client
        clients.forEach((elem) => {
            if (elem.id === meetings[key].client) {
                client = elem.company
            }
        });

        list.innerHTML += `
                <li class="meet_list_elem">                                
                <div class="meet_list_elem_data">${meetings[key].date}</div>
                <div class="meet_list_elem_time">${str}</div>
                <div class="meet_list_elem_time1">${str1}</div>
                <div class="meet_list_elem_place">${place}</div>
                <div class="meet_list_elem_theme">${meetings[key].topic}</div> 
                <div class="meet_list_elem_klient">${client}</div>
            </li>`
    }

}