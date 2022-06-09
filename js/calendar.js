window.onload = async function () {
    let book_room = document.getElementById('book_room')
    let save_button = document.getElementById('savethisshit1')
    let place = document.getElementById('place_list')

    let responce = await fetch('http://176.119.157.82:8000/api/meeting?booking=true')

    let booking  = await responce.json()

    let meet_list_list1 = document.querySelector('.meet_list_list')


    let keyS



    for(keyS in booking) {

        meet_list_list1.innerHTML += `
        <li class="meet_list_elem">
            <div class="meet_list_elem_theme">${booking[keyS].creatorText}</div> 
            <div class="meet_list_elem_place">Переговорка 1.${booking[keyS].place}</div>
            <div class="meet_list_elem_data">${booking[keyS].date}</div>     
            <div class="meet_list_elem_time">${booking[keyS].start}</div>
            <div class="meet_list_elem_time1">${booking[keyS].end}</div>  
        </li>

        `
    }


    let mark = 1

    book_room.onclick = async function() {
        let responce = await fetch('http://176.119.157.82:8000/api/place/')

        let places = await responce.json()

        let list_place = document.querySelector('.place_input')

        let keyS

        if (mark == 1) {
            for (keyS in places) {
                list_place.innerHTML += `
                
                <option value="${places[keyS].name}" id="qwe">${places[keyS].name}</option>           `
            }
        }

        mark = 0
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
        let data = document.getElementById("data")
        let creator = document.getElementById("client")

        let body = {

            date: data.value,
            start: time_start.value,
            end: time_end.value,
            topic: "",
            creator: 1,
            place: id_place,
            client: null,
            contact: null,
            creatorText: creator.value

        }

        let responce = await sendRequest('POST', 'http://176.119.157.82:8000/api/meeting/', body)

        // let empl = await responce.json()


        // let id_clien = empl.creator
        // let id_meet = empl.id

        // fetch('http://176.119.157.82:8000/api/employeelist/', {
        //     method: 'POST',
        //     body: JSON.stringify({
        //         "employee": id_clien,
        //         "meeting": id_meet
        //     }),
        //     headers: {
        //         'Content-type': 'application/json; charset=UTF-8',
        //     },
        // })

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

        window.location = 'http://sbermeeting.tk/calendar.html'

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

}

async function get_first_id(api_path) {
    let responce = await fetch('http://176.119.157.82:8000/api/' + api_path)
    let objects = await responce.json()
    return objects[0].id
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