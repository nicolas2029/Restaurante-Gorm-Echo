function loadRow(data) {
    let temp = 
    `<div class="col-lg-6 menu-item">
    <div class="menu-content">
    <a href="#">` + data.name + `</a><span>$`+ data.price + `</span>
    </div>
    <div class="menu-ingredients">` + data.description + `</div>
    </div>`;
    return temp
    //document.getElementById('body').innerHTML = temp;
}

function getTable(url){
    let myInit = { method: 'GET',
            };
    let myRequest = new Request(url, myInit);
    fetch(myRequest)
    .then(res => {
        res.json().then(
        data => {
            if (data.length > 0) {
            let temp = "";
            data.forEach((itemData) => {
                temp = loadRow(itemData)
            });
            document.getElementById('menu-container').innerHTML = temp;
            }
        }
        )
    }
    )
}

function getTa(url){
    var myInit = { method: 'GET',
            };
    var myRequest = new Request(url, myInit);
    fetch(myRequest)
    .then(res => {
        res.json().then(
        data => {
            if (data.length > 0) {
            var temp = "";
            data.forEach((itemData) => {
                temp += "<tr>";
                temp += "<td>" + itemData.name + "</td>";
                temp += "<td>" + itemData.price + "</td>";
                temp += "<td>" + itemData.description + "</td></tr>";
            });
            document.getElementById('body').innerHTML = temp;
            }
        }
        )
    }
    )
}
