
let myAdmin;
let nameImage;
let myRole;
let roles;
let establishments;
let mapProduct = new Map();
let myMap = new Map();
let myEstablishment;
let total = 0
let ordersPending = new Map(); //Pendiente a ver si lo utilizare
let mapOrderTableTable = new Map();
let mapTableOrder = new Map();
let mapOrderTable = new Map();
let modal;

function closeModal(){
    modal.innerHTML = "";
    modal.style.display = "none";
}

function openModal(funcName, action){
    modal.innerHTML = `<div class="modal-content">
        <header class="close" id="modal-close" onclick="closeModal()">&times;</header>
        <div class="section-title">
        <h1 style="font-weight: bold">Confirmar<span style="color: #ffb03b"> ${action}</span></h1>
        <p>¿Estás seguro de continuar?</p>
        </div>
        <div>
        <button type="button" onclick="closeModal()" class="cancelbtn">Cancelar</button>
        <button type="button" onclick="${funcName}()" class="deletebtn">Confirmar</button>
        </div>
    </div>`
    modal.style.display = "block";
}

window.onclick = function(event) {
    if (event.target == modal) {
        modal.style.display = "none";
    }
}


const PERMISSIONS = {
    WhitOutRestriction:1,

}

function deleteElementByID(id){
	element = document.getElementById(id);	
	if (!element){
		alert("El elemento selecionado no existe");
	} else {
		parent = element.parentNode;
		parent.removeChild(element);
	}
}

function addOrderToPending(o) {
    document.getElementById("row-order").innerHTML += loadOrder(o, 0);
}

function addOrderProductToOrder(tableID, products) {
    let orderProduct = ""
    products.forEach(val => {orderProduct += loadOrderProduct(val)});
    document.getElementById(`box-pending-${tableID}`).innerHTML += orderProduct;
}

function addProductsToOrder(tableID, products){
    console.log(tableID)
    let orderID = mapTableOrder.get(tableID)
    fetch(`http://localhost:80/api/v1/order/${orderID}`, {
        method:"POST",
        headers:{'Content-Type': 'application/json'},
        body:JSON.stringify(products)
    }).then(res => {
        if (res.ok){
            addOrderProductToOrder(tableID, products)
        }
    })
}

function loadOrderProduct(product) {
    total += product.amount*mapProduct.get(product.product_id).price
    return `<h5>${mapProduct.get(product.product_id).name}</h5>
    <p>Precio unitatio - ${mapProduct.get(product.product_id).price.toFixed(2)}</p>
    <p>Cantidad - ${product.amount}</p>
    <p>Precio total - ${product.amount*mapProduct.get(product.product_id).price.toFixed(2)}</p>`
}

function completeOrder(orderID) {
    fetch(`http://localhost:80/api/v1/order/${orderID}`, {
        method:"POST",
        headers:{'Content-Type': 'application/json'},
    }).then(res => {
        if (res.ok){
            deleteElementByID(`order-pending-${orderID}`);
            mapTableOrder.delete(mapOrderTable.get(orderID))
            mapOrderTable.delete(orderID)
        }})
}

function loadOrder(op, i) {                             //
    let order; 
    let orderProduct = ``
    //addOrderToPending(op);
    total=0;
    let tableID = mapOrderTableTable.get(op.order.table_id);
    //ordersPending.set(op.order, op) // agrega al map de ordenes pendientes
    mapTableOrder.set(tableID, op.order.id)
    mapOrderTable.set(op.order.id, tableID)
    op.order_products.forEach(val => {orderProduct += loadOrderProduct(val)});
    order = `<div class="col-lg-3" id="order-pending-${op.order.id}">
    <div class="box">
    <div id="box-pending-${tableID}">
        <span>Mesa: ${tableID}</span>
        <h4>Encabezado</h4>
        <p>Fecha: ${op.order.created_at}</p>
        <p>Precio Total: ${total.toFixed(2)}</p>
        <h4>Productos</h4>`
    order += orderProduct
    order += `</div><div class="text-center"><button type="button" onclick="completeOrder(${op.order.id})">Completar pedido</button></div></div></div>`;
    return order;
}

function loadOrderSection(o) {
    let temp = 
    `<div class="container">

    <div class="section-title">
        <h2>Mis Pedidos <span>Pendientes</span></h2>
        <p>Aqui podras ver todos los pedidos sin completar que has realizado</p>
    </div>
    <div class="row" id="row-order">`;
    if (o != null){
        o.forEach((op,i) => {temp += loadOrder(op, i)});
    }
    temp += `</div></div>`;
    document.getElementById("my-incomplete-orders").innerHTML=temp
}


function updateTotal(amount, price, id, item, key) {
    if (amount > 0){
        document.getElementById(id).innerText= `Precio Total: ${amount*price}`;
        myMap.set(key, amount);
    }else{
        let node = document.getElementById(item);
        node.parentNode.removeChild(node);
        myMap.delete(key);
    }
}

function showSelectEstablishments() {
    let temp = `<option value="" selected disabled hidden>Selecciona un establecimiento</option>`;
    establishments.forEach(a => {
        temp += `<option value="${a.id}" >${a.id}</option>`;
    });
    document.getElementById("hire-select-establishment").innerHTML = temp;
    temp += `<option value="0" selected >Nuevo Establecimiento</option>`
    document.getElementById("crud-select-establishment").innerHTML = temp;
}

/*function SetEstablishments(a) {
    establishments = a;
    showHireUser()
    showFireUser()
}*/

function loadAllEstablishments(){
    fetch(`http://localhost:80/api/v1/establishment/`,{method:"GET"}).then(res => res.json().then(data => {
        establishments = data;

        showCRUDEstablishment();
        showHireFireAdmin();
    })).catch(a => console.log(a));
}
function SetMyAdmin(a) {
    myAdmin = a;
    showMenu();
    getMyRolePermissions();
}

function loadMyEstablishments(){
    fetch(`http://localhost:80/api/v1/establishment/${myAdmin.establishment_id}`,{method:"GET"}).then(res => res.json().then(data => {
        myEstablishment = data;
        showMyOrder();
    })).catch(a => console.log(a));
}

/*function SetAllRoles(a) {
    roles = a;
    
}*/

function SetMyRole(a) {
    myRole = a;
    showAllFunctionsByPermissions();
}

function getMyRolePermissions(){
    fetch(`http://localhost:80/api/v1/rol/${myAdmin.rol_id}`,{method:"GET"}).then(res => res.json().then(data => {
        console.log("rol: ",data);
        SetMyRole(data);
    })).catch(a => console.log(a));
}

function getAllRoles(per){
    console.log(roles)
    if (roles == null){
        fetch(`http://localhost:80/api/v1/rol/`,{method:"GET"}).then(res => res.json().then(data => {
        roles = data;
        if (per == 2){
            loadAllEstablishments();
        }
        if (per == 16){
            showHireFireUser();
        }
        })).catch(a => console.log(a));
    }else{
        if (per == 3){
            loadAllEstablishments();
        }
        if (per == 16){
            showHireFireUser();
        }
    }
    
}

function setImageToProduct(id, name) {
    let file = document.getElementById(id);
    fetch(`http://localhost:80/api/v1/product/img/${name}`,{
        method:"POST",
        body:file,
        headers:{
            'Content-Type': 'multipart/form-data'
        }
    }).then(console.log("IMAGEN AGREGADA CON EXITO"));

}

function updateUserRol(){
    let id = parseInt(document.getElementById("hire-select-rol").value);
    let mail = document.getElementById("hire-input-email").value;
    fetch(`http://localhost:80/api/v1/user/${mail}`,{
        method:"PUT",
        body:{
            "rol_id":id,
        },
        headers:{
            'Content-Type': 'multipart/form-data'
        }
    }).then(console.log("IMAGEN AGREGADA CON EXITO"));
}

function getRolesCanBeSet(){
    let rolesCanBeSet = `<option value="" selected disabled hidden>Selecciona un rol</option>`;
    let i = myAdmin.rol_id;
    for (;i < roles.length; i++){
        rolesCanBeSet += `<option value="${roles[i].id}" >${roles[i].name}</option>`;
    }
    return rolesCanBeSet;
}

function hireUser(){
    let rol_id = parseInt(document.getElementById("hire-select-rol").value);
    let establishment_id = parseInt(document.getElementById("hire-select-establishment").value);
    let email = document.getElementById("hire-input-email").value;
    fetch(`http://localhost:80/api/v1/admin/hire/${email}`,{
        method:"PATCH",
        headers:{
            'Content-Type': 'application/json'
        },
        body:JSON.stringify({
            "rol_id":rol_id,
            "establishment_id":establishment_id,
        })
    })
}

function hireUserInStablishment(){
    let rol_id = parseInt(document.getElementById("hire-select-rol-st").value);
    let email = document.getElementById("hire-input-email-st").value;
    fetch(`http://localhost:80/api/v1/user/hire/${email}`,{
        method:"PATCH",
        headers:{
            'Content-Type': 'application/json'
        },
        body:JSON.stringify({
            "rol_id":rol_id
        })
    })
}


function crudEstablishment(){
    let address = document.getElementsByName("form-crud-establishment")
    let dataAddress = {
        "line1":address.item(0).value,
        "line2":address.item(1).value,
        "city":address.item(2).value,
        "state":address.item(3).value,
        "country":address.item(4).value,
        "postal_code":address.item(5).value
    };
    let amountTables = address.item(6).value;
    let st = address.item(7).value;
    let myInit = {method:"POST",headers:{'Content-Type': 'application/json'}};
    let myRequest;
    let addressInit={
        method:"",
        body:JSON.stringify(dataAddress),
        headers:{'Content-Type': 'application/json'}
    };
    if (st == "0"){
        addressInit.method = "POST";
        fetch("http://localhost:80/api/v1/address/", addressInit).then(res => res.json().then(data =>{
            console.log(data);
            myInit.body = JSON.stringify({address_id:data});
            fetch(`http://localhost:80/api/v1/establishment/${amountTables}`, myInit).then(res => console.log(res));
        }))
    }else{
        addressInit.method = "PUT";
        fecht(`http://localhost:80/api/v1/address/${st}`, addressInit).then(res => console.log(res));
    }
}

function addProduct(key, amount, src, name, price, description) {
    let r = myMap.get(key);
    let idInput = `input-product-${key}`;
    if (r === undefined) {
        let id=`total-product-${key}`;
        let idItem = `product-${key}`;
        let temp = 
        `<div class="col-lg-2 col-md-3" id=${idItem}>
            <div>
                <div class="gallery-item">
                    <img src="${src}" alt="${description}" class="img-fluid">
                </div>
                <div align="center">
                    <label>${name}</label>
                </div>
                <div>
                    <input class="form-control" id="${idInput}" type="number" min=0 onchange="updateTotal(this.value, ${price}, '${id}', '${idItem}', ${key})">
                </div>
                <div align="center">
                    <label id="${id}">Precio Total: ${price}</label>
                </div>
            </div>
        </div>`;
        myMap.set(key, amount);
        document.getElementById("shopping-cart").innerHTML += temp
        myMap.forEach((val,keyMap) =>{
        document.getElementById(`input-product-${keyMap}`).value=val
    })
    }
    
}

function loadMenuProduct(data) {
    var temp = 
    `<div class="col-lg-6 menu-item">
    <div class="menu-content">
    <button name="button-menu" onclick="addProduct(${data.id}, 1, '${data.img}', '${data.name}', ${data.price}, '${data.description}')">${data.name}</button><span>$${data.price}</span>
    </div>
    <div class="menu-ingredients">${data.description}</div>
    </div>`;
    mapProduct.set(data.id, {
        name:`${data.name}`,
        price:data.price
    });
    return temp;
}


function fireUser(){
    let email = document.getElementById("fire-input-email").value;
    fetch(`http://localhost:80/api/v1/admin/fire/${email}`,{
        method:"PATCH",
    })
}

function fireUserInStablishment(){
    let email = document.getElementById("fire-input-email-st").value;
    fetch(`http://localhost:80/api/v1/user/fire/${email}`,{
        method:"PATCH",
    })
}

function menuEmployee(data){
    let temp = "";
    data.forEach(item => temp+=loadMenuProduct(item)); 
    document.getElementById("menu").innerHTML=`<div class="container">
    <div class="section-title">
        <h2>Revisa nuestro sabroso <span>Menú</span></h2>
    </div>
    <div class="row menu-container" id="menu-container">
    ${temp}
    </div>

    </div>`;
}

function menuAdmin(data){
    let temp = "";
    data.forEach(item => temp+=loadMenuProduct(item)); 
    document.getElementById("menu").innerHTML=`<div class="container">
    <div class="section-title">
        <h2>Revisa nuestro sabroso <span>Menú</span></h2>
    </div>
    <div class="row menu-container" id="menu-container">
    ${temp}
    </div>
    </div>`;
}

function showMenu(){
    fetch(`http://localhost:80/api/v1/product/`).then(res => res.json().then(data => {
        let temp = "";
        data.forEach(item => temp+=loadMenuProduct(item)); 
        document.getElementById("menu").innerHTML=`<div class="container">
        <div class="section-title">
            <h2>Revisa nuestro sabroso <span>Menú</span></h2>
        </div>
        <div class="row menu-container" id="menu-container">
        ${temp}
        </div>

        </div>`;
        //document.getElementById("menu-container").innerHTML = temp;
    }))
}

function getSelectTableByEstablishment(){
    let count = 1;
    let temp =`"<option value="" selected disabled hidden>Selecciona una mesa</option>"`;
    myEstablishment.tables.forEach(data =>{
        console.log(data)
        temp += `<option value="${data.id}" >${count}</option>`;
        mapOrderTableTable.set(data.id, count);
        count +=1;
    })
    return temp
}

function getAllOrdersPendingByUser(){
    fetch("http://localhost:80/api/v1/order/user/").then(res =>
        {
            if (res.status == 200){
                res.json().then(data => {
                    loadOrderSection(data)
                })
            }else if(res.status == 204){
                loadOrderSection(null)
            }
        })
}

function switchOrderAction(){
    addOrderProductToOrder(tableID, op)
}

function makeOrderInEstablishment(){
    let t = parseInt(document.getElementById("select-table").value);
    let products = [];
    myMap.forEach((val, key) => {
        let p = {
            "product_id":parseInt(key),
            "amount":parseInt(val)
        };
        products.push(p);
    });
    //console.log(t);
    if(mapTableOrder.has(mapOrderTableTable.get(t))){
        addProductsToOrder(mapOrderTableTable.get(t), products)
    }else{
        let orderProduct = {
            "order":{
                "establishment_id":myAdmin.establishment_id,                                   //CAMBIAR LO DE LOS ESTABLECIMIENTOS EN AMBOS TIPOS DE ORDENES QUEDARA PENDIENTE.
                "table_id":t
            },
            "order_products":products
        };
        fetch("http://localhost:80/api/v1/order/", {
            //credentials: 'same-origin',
            method: 'POST',
            body: JSON.stringify(orderProduct),
            headers:{
                'Content-Type': 'application/json'
            }}).then(res => {
                document.getElementById("shopping-cart").innerHTML = ""
                myMap.clear()
                document.getElementById("select-table").value = ""
                res.json().then(data => addOrderToPending(data))})// Limpiar formulario, agregar a ordenes pendientes
    }
    }

function showMyOrder(){
    document.getElementById("my-order").innerHTML = `<div class="container-fluid">
    <div class="section-title">
        <h2>Nuevo <span>Pedido</span></h2>
        <p>Aqui se listaran todos los productos seleccionados para proceder con su respectivo pedido</p>
    </div>
    <div class="container">
        <div id="shopping-cart" class="row no-gutters"></div>
        <form class="php-email-form" action="">
        <div class="form-row">
            <select name="form-my-order" id="select-table" required>
            </select>
            </div>
            <div class="mb-3">
            </div>
            <div class="text-center"><button type="button" onclick="openModal('makeOrderInEstablishment', 'Pedido')">Realizar pedido</button></div>
        </div>
        </form>
    </div>`;
    document.getElementById("select-table").innerHTML = getSelectTableByEstablishment();
    getAllOrdersPendingByUser();                        //
}

function showHireFireAdmin(){
    showHireUser();
    showFireUser();
}
function showHireFireUser(){
    showHireUserInStablishment();
    showFireUserInStablishment();
}

function showCRUDEstablishment(){
    document.getElementById("create-establishment").innerHTML = `<div class="container">
    <div class="section-title">
    <h2>Registrar <span>Establecimiento</span></h2>
    <p>Agrega un nuevo establecimiento; pendiente por actualizar</p>
    </div>
    </div>
    <div class="container book-a-table">
    <form class="php-email-form" action="" >
    <div class="form-row">
        <div class="col-lg-4 col-md-6 form-group">
        <input name="form-crud-establishment" type="text" class="form-control" id="line1" placeholder="Ingresa la Calle y numero del local" required>
        </div>
        <div class="col-lg-4 col-md-6 form-group">
        <input name="form-crud-establishment" type="text" class="form-control" id="line2" placeholder="Ingresa la Colonia">
        </div>
        <div class="col-lg-4 col-md-6 form-group">
        <input name="form-crud-establishment" type="text" class="form-control" id="city" placeholder="Ingresa la Ciudad" required>
        </div>
        <div class="col-lg-4 col-md-6 form-group">
        <input name="form-crud-establishment" type="text" class="form-control" id="state" placeholder="Ingresa el Estado" required>
        </div>
        <div class="col-lg-4 col-md-6 form-group">
        <input name="form-crud-establishment" type="text" class="form-control" id="country" placeholder="Ingresa el Pais" value="Mexico" required>
        </div>
        <div class="col-lg-4 col-md-6 form-group">
        <input name="form-crud-establishment" type="text" class="form-control" id="postal_code" placeholder="Codigo Postal" required>
        </div>
        <div class="col-lg-4 col-md-6 form-group">
        <input name="form-crud-establishment" type="number" class="form-control" id="amount-table" placeholder="Ingresa la cantidad de mesas" required>
        </div>
        <div class="col-md-6 form-group">
        <select name="form-crud-establishment" id="crud-select-establishment" required>
        </select>
    </div>
    <div class="text-center" onclick="openModal('crudEstablishment', 'Establecimiento')"><button type="button">Registrar establecimiento</button></div>
    </form>
    </div>`
}

function showFireUser(){
    document.getElementById("fire-employee").innerHTML= `<div class="container">
        <div class="section-title">
        <h2>Despedir <span>Empleado</span></h2>
        <p>Formulario para despedir un empleado</p>
        </div>
    </div>
    <div class="container book-a-table">
    <form class="php-email-form" action="">
        <div class="form-row">
        <div class="col-md-6 form-group">
            <input type="email" name="fire-form" class="form-control" id="fire-input-email" placeholder="Email">
        </div>
        <div class="text-center"><button type="button" onclick="openModal('fireUser', 'Despido')">Despedir Empleado</button></div>
    </form>
    </div>`;
}

function showFireUserInStablishment(){
    document.getElementById("fire-employee-st").innerHTML= `<div class="container">
        <div class="section-title">
        <h2>Despedir <span>Empleado</span></h2>
        <p>Formulario para despedir un empleado</p>
        </div>
    </div>
    <div class="container book-a-table">
    <form class="php-email-form" action="">
        <div class="form-row">
        <div class="col-md-6 form-group">
            <input type="email" name="fire-form-st" class="form-control" id="fire-input-email-st" placeholder="Email">
        </div>
        <div class="text-center"><button type="button" onclick="openModal('fireUserInStablishment', 'Despido')">Despedir Empleado</button></div>
    </form>
    </div>`;
}

function showHireUserInStablishment(){
    document.getElementById("hire-employee-st").innerHTML= `<div class="container">
        <div class="section-title">
        <h2>Contratar <span>Empleado</span></h2>
        <p>Formulario para dar de alta a un usuario dentro de un establecimiento</p>
        </div>
    </div>
    <div class="container book-a-table">
    <form class="php-email-form" action="">
        <div class="form-row">
        <div class="col-md-6 form-group">
            <input type="email" name="hire-form-st" class="form-control" id="hire-input-email-st" placeholder="Email">
        </div>
        <div class="col-md-6 form-group">
        <select name="hire-form-st" id="hire-select-rol-st" required>
        </select>
        </div>
        <div class="text-center"><button type="button" onclick="openModal('hireUserInStablishment', 'Contratación')">Contratar Empleado</button></div>
    </form>
    </div>`;
    document.getElementById("hire-select-rol-st").innerHTML = getRolesCanBeSet();
}

function showHireUser(){
    document.getElementById("hire-employee").innerHTML= `<div class="container">
        <div class="section-title">
        <h2>Contratar <span>Empleado</span></h2>
        <p>Formulario para dar de alta a un usuario dentro de un establecimiento</p>
        </div>
    </div>
    <div class="container book-a-table">
    <form class="php-email-form" action="">
        <div class="form-row">
        <div class="col-md-6 form-group">
            <input type="email" name="hire-form" class="form-control" id="hire-input-email" placeholder="Email">
        </div>
        <div class="col-md-6 form-group">
        <select name="hire-form" id="hire-select-rol" required>
        </select>
        </div>
        <div class="col-md-6 form-group">
        <select name="hire-form" id="hire-select-establishment" required>
        </select>
        </div>
        <div class="text-center"><button type="button" onclick="openModal('hireUser', 'Contratación')">Contratar Empleado</button></div>
    </form>
    </div>`;
    document.getElementById("hire-select-rol").innerHTML = getRolesCanBeSet();
    showSelectEstablishments()

}

function createProduct() {
    let name = document.getElementById("input-name").value;
    let price = parseFloat(document.getElementById("input-price").value);
    let description = document.getElementById("input-description").value;
    var file = document.getElementById("input-file").files[0];
    if (name == null || price == null || description == null || file == null){
        return false
    }
    fetch("http://localhost:80/api/v1/product/",{
        method: "POST",
        body: JSON.stringify({
            "name":name,
            "price":price,
            "description":description,
            "img":`assets/img/products/${name}.jpg`
        }),
        headers:{
            'Content-Type': 'application/json'
        }
    }).then(res => {
        if (res.ok){
            sendFormSetImage();
            
        }else{
                console.log(res.status);
            }
        $('#form-set-image').trigger("reset");
        document.getElementById("img") =`<img id="img" class="img-fluid">`
        });
}

async function isLogin (){
    var myInit = {method: 'GET'};
    var myRequest = new Request('http://localhost:80/api/v1/user/login/', myInit);
    fetch(myRequest).then(res => {
        if (res.ok){
            res.json().then(a =>{
                SetMyAdmin(a);
            })
        }
    })
}

function updateNameImage(){
    nameImage = document.getElementById("input-name").value
    document.getElementById("submit-product").formAction = `http://localhost:80/api/v1/product/img/${nameImage}`
    console.log(nameImage)
}



function sendFormSetImage(){
    nameImage = document.getElementById("input-name").value;
    document.getElementById("form-set-image").action = `http://localhost:80/api/v1/product/img/${nameImage}`;
    document.getElementById("form-set-image").submit();
    
}

$('#form-set-image').submit(function () {
    $.ajax({
        type: $('#form-set-image').attr('method'), 
        url: `"http://localhost:80/api/v1/product/img/${nameImage}"`,
        data: $('#form-set-image').serialize(),
        success: function (data) { console.log("Enviado") }  
    });
    return false;
});

function showCreateProduct(){
    document.getElementById("create-product").innerHTML=`<div class="container">
        <div class="section-title">
        <h2>Registrar <span>Productos</span></h2>
        <p>Registro de productos nuevos</p>
    </div>
    </div>
    <div class="container book-a-table">
    <form class="php-email-form" method="post" enctype="multipart/form-data" id="form-set-image" target="request">
        <div class="form-row" >
        <div class="col-md-6 form-group" id="div-input-name">
            <input type="text" name="name" class="form-control" id="input-name" oninput="updateNameImage()" placeholder="Nombre del producto">
        </div>
        <div class="col-md-6 form-group">
            <input type="number" class="form-control" name="price" id="input-price" placeholder="Precio" >
        </div>
        <div class="col-md-6 form-group">
            <input class="form-control" name="file" type="file" id="input-file" accept="image/*" onchange="mostrar()" >
        </div>
        <div class="gallery-item">
            <img id="img" class="img-fluid">
        </div>
        </div>
        <div class="form-group">
        <textarea id="input-description" class="form-control" name="message" rows="5" data-rule="required" data-msg="Please write something for us" placeholder="Descripción"></textarea>
        </div>
        <div class="text-center"><button type="button" onclick="openModal('createProduct', 'Nuevo Producto')">Registrar producto</button></div>
    </form>
    </div>
    
    <iframe name="request" width="0" height="0" frameborder="0" id="request" style="display: none;"></iframe>`
    
}

function showAll(){
    showCreateProduct();
    getAllRoles(2);
    //showMyOrder();
}

function showMakerOrderRemote(){

}

function showFunctionByPermission(id){
    console.log(id)
    switch (id) {
        case 2:
            showCreateEstablishment();
            break
        case 3:
            getAllRoles(2);
            break;
        case 4:
            
            break;
        case 5:
            loadMyEstablishments();
            break;
        case 6:
            break;
        case 7:
            break;
        case 8:
            break;
        case 9:
            break;
        case 10:
            break;
        case 11:
            break;
        case 12:
            break;
        case 13:
            break;
        case 14:
            break;
        case 15:
            break;
        case 16:
            getAllRoles(16);
            break;
        default:
            break;
    }
}

function showAllFunctionsByPermissions(){
    console.log(myRole);
    if(myRole.id == 1){
        showAll()
    }else{
        myRole.permissions.forEach(permission => {
            showFunctionByPermission(permission.ID)
        })
    }
}

function loadAdminPage() {
    modal = document.getElementById("myModal");
    isLogin();
}

function mostrar(){
    var file = document.getElementById("input-file").files[0];
    var reader = new FileReader();
    if (file) {
        reader.readAsDataURL(file );
        reader.onloadend = function () {
            document.getElementById("img").src = reader.result;
        }
        }
}