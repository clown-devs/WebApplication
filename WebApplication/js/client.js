window.onload = async function () {


    let responce = await fetch('http://176.119.157.82:8000/api/client/')
    let content = await responce.json()
    let list = document.querySelector('.back')

    let key;

    for (key in content) {

        list.innerHTML += `
           <li class="klient">
           <img src="./img/klient_ava.png" class="klient_ava">
           <div class="klient_name">${content[key].company}</div>
           <div class="klient_inn">${content[key].inn}</div>
       </li>
           `


    }
    
}