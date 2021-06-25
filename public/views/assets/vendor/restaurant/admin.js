
let myAdmin;
let nameImage;
let nameUpdateImage;
let myRole;
let roles;
let mapProduct = new Map();
let mapProductUpdate = new Map();
let myMap = new Map();
let myEstablishment;
let total = 0
let ordersPending = new Map(); //Pendiente a ver si lo utilizare
let mapOrderTableTable = new Map();
let mapTableOrder = new Map();
let mapOrderTable = new Map();
let payments;

function updatePrice(id){
    document.getElementById(`p-${id}`).innerHTML = "Precio Total: "
}

function showCrudPayments(){
    let temp = ``;
    payments.forEach(data => {temp += `<option id="op-pay-${data.id}" value="${data.id}">${data.name}</option>`;
});
    document.getElementById("crud-payment").innerHTML = `<div class="container">
    <div class="section-title">
    <h2>Crear / Actualizar <span>métodos de pago</span></h2>
    <p>Formulario para crear o modificar métodos de pago</p>
    </div>
    </div>
    <div class="container book-a-table">
    <form class="php-email-form" action="">
    <div class="form-row">
    <div class="col-md-6 form-group">
        <input type="text" name="pay-form" class="form-control" id="pay-name" placeholder="Nombre">
    </div>
    <div class="col-md-6 form-group">
        <select name="pay-form" id="pay-select" required=""><option value="0" selected="">Nuevo</option>${temp}</select>
    </div>
    <div class="text-center"><button id="pay-button" type="button" onclick="openModal('crudPayment', 'Proceso')">Enviar</button></div>

    </div></form></div>`;
    document.getElementById("nav-crud-payment").innerHTML = `<a href="#crud-payment">métodos de pago</a>`
}

function crudPayment(){
    let name = document.getElementById("pay-name").value;
    let id = parseInt(document.getElementById("pay-select").value);
    let url;
    let method;
    if (name == ""){
        return "Necesitas ingresar el nombre del método de pago";
    }
    if (id == 0){
        url = `http://localhost:80/api/v1/pay/`;
        method = `POST`;
    }else if (id > 0){
        url = `http://localhost:80/api/v1/pay/${id}`;
        method = `PUT`;
    }else{
        return "ID no válida...";
    }
    fetch(url, {
        method:method,
        headers:{'Content-Type': 'application/json'},
        body:JSON.stringify({name:name})
    }).then(res => {
        if (res.status == 201){
            res.json().then(obj => document.getElementById("pay-select").innerHTML += `<option value="${obj.id}" id="op-pay-${obj.id}">${name}</option>`);
            showSuccessful();
        }else if (res.ok){
            document.getElementById(`op-pay-${id}`).innerHTML = name;
            showSuccessful();
        }else {
            res.json().then(x => showError(translateError(x.message)));
        }
        document.getElementById("pay-name").value = "";
        document.getElementById("pay-select").value = "0";
    });
}

function fetchPayments(){
    fetch(`http://localhost:80/api/v1/pay/`).then(res => {
        if (res.ok) {
            res.json().then(data => {payments = data; 
                showCrudPayments();
            });
        }
    })
}

function timeToString(t){
    time = t
    let year = time.slice(0,4)
    let month = time.slice(5,7)
    let day = time.slice(8, 10)
    let hour = time.slice(11,13)
    let minute = time.slice(14,16)
    return `${day}/${month}/${year} ${hour}:${minute}`;
}

function getShowOrderPendingByEstablishment(op){
    let order; 
    let orderProduct = ``
    //addOrderToPending(op);
    total=0;
    let button = "";
    let tableID = mapOrderTableTable.get(op.order.table_id);
    console.log("ORDENES: ",op);
    if(!mapOrders.has(op.order.id)){
        mapOrders.set(op.order.id, op);
    }
    //ordersPending.set(op.order, op) // agrega al map de ordenes pendientes
    mapTableOrder.set(tableID, op.order.id)
    mapOrderTable.set(op.order.id, tableID)
    op.order.products.forEach(val => {
        if(!mapProductUpdate.has(val.id)){
            mapProductUpdate.set(val.id, val)
        }
    })
    op.order_products.forEach(val => {orderProduct += loadOrderProduct(val)});
    if(op.order.table_id == null){
        button = `<button type="button" onclick="openModalParam('completeOrderRemote', 'Pedido', ${op.order.id})">Completar pedido</button>`;
    }
    order = `<div class="col-lg-3" id="order-pending-e-${op.order.id}">
    <div class="box">
    <div id="box-pending-e-${op.order.id}">
        <h4>Encabezado</h4>
        <p>Fecha: ${timeToString(op.order.created_at)}</p>
        <p>Precio Total: ${total.toFixed(2)}</p>
        <h4>Productos</h4>`;
    order += orderProduct;
    order += `</div><div class="text-center">${button}</div></div></div>`;
    return order;
}


function getAllOrdersPendingByEstablishment(){
    fetch(`http://localhost:80/api/v1/order/establishment/`).then(res => {
        if (res.ok){ 
            let temp=`<div class="container">
                <div class="section-title">
                    <h2>Pedidos <span>Pendientes</span></h2>
                    <p>Aquí podras ver todos los pedidos sin completar</p>
                </div>
                <div class="row">`;
            if(res.status == 204){
                temp +=`</div></div>`
                document.getElementById("order-pending-e").innerHTML = temp;
            }else{
                res.json().then(data => {
                if(data != null){
                    data.forEach(val => {
                        temp += getShowOrderPendingByEstablishment(val);
                    })
                    temp +=`</div></div>`;
                    document.getElementById("order-pending-e").innerHTML = temp;
                }
                
            })
            }
            
        }else{
        }
        document.getElementById("nav-order-pending").innerHTML = `<a href="#order-pending-e">Pedidos Pendientes</a>`
    })
}

function getShowOrder(op){
    let order; 
    let orderProduct = ``
    //addOrderToPending(op);
    total=0;
    let tempPay = ``;
    let tableID = mapOrderTableTable.get(op.order.table_id);
    console.log(op);
    if(!mapOrders.has(op.order.id)){
        mapOrders.set(op.order.id, op);
    }
    //ordersPending.set(op.order, op) // agrega al map de ordenes pendientes
    mapTableOrder.set(tableID, op.order.id)
    mapOrderTable.set(op.order.id, tableID)
    op.order.products.forEach(val => {
        if(!mapProductUpdate.has(val.id)){
            mapProductUpdate.set(val.id, val)
        }
    })
    op.order_products.forEach(val => {orderProduct += loadOrderProduct(val)});
    if(op.order.pay_id != null && op.order.pay_id != 0){
        tempPay = `<p>Tipo de pago: ${mapPayments.get(op.order.pay_id)}</p>`;
    }
    order = `<div class="col-lg-3" id="order-pending-e-${op.order.id}" onclick="printPDF(${op.order.id})">
    <div class="box">
    <div id="box-pending-e-${op.order.id}">
        <h4>Encabezado</h4>
        <p>Fecha: ${timeToString(op.order.created_at)}</p>
        ${tempPay}
        <p>Precio Total: ${total.toFixed(2)}</p>
        <h4>Productos</h4>`
    order += orderProduct
    order += `</div></div></div>`;
    return order;
}

function getAllOrdersByEstablishment(){
    fetch(`http://localhost:80/api/v1/order/establishment/all`).then(res => {
    if (res.ok){
        res.json().then(data => {
            console.log(data)
            let temp=`<div class="container">
            <div class="section-title">
                <h2>Pedidos <span></span></h2>
                <p>Aquí podras ver todos los pedidos</p>
            </div>
            <div class="row">`;
            if(data != null){
                data.forEach(val => {
                    temp += getShowOrder(val);
                })
            }
            temp +=`</div></div>`
            document.getElementById("orders-by-stablishment").innerHTML = temp;
            document.getElementById("nav-order-st").innerHTML = `<a href="#orders-by-stablishment">Pedidos</a>`
        })
    }else{
    }
    })
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
    console.log(o);
    document.getElementById("row-order").innerHTML += loadOrder(o, 0);
}

function addOrderProductToOrder(tableID, products) {
    let orderProduct = ""
    total = 0;
    products.forEach(val => {orderProduct += loadOrderProduct(val)});
    total += parseInt(document.getElementById(`i-p-${tableID}`).value);
    document.getElementById(`i-p-${tableID}`).value = total;
    document.getElementById(`p-${tableID}`).innerHTML = `Precio Total: ${total.toFixed(2)}`;
    document.getElementById(`box-pending-${tableID}`).innerHTML += orderProduct;
}

function addProductsToOrder(tableID, products){
    if (!mapTableOrder.has(tableID)){return `La mesa no esta en uso`;}
    let orderID = mapTableOrder.get(tableID)
    fetch(`http://localhost:80/api/v1/order/${orderID}`, {
        method:"PUT",
        headers:{'Content-Type': 'application/json'},
        body:JSON.stringify(products)
    }).then(res => {
        document.getElementById("shopping-cart").innerHTML = "";
        myMap.clear();
        document.getElementById("select-table").value = "";
        if (res.ok){
            showSuccessful();
            addOrderProductToOrder(tableID, products);
        }else{
            res.json(x => showError(translateError(x.message)));
        }
    })
}

function loadOrderProduct(product) {
    if(mapProduct.has(product.product_id)){
        total += product.amount*mapProduct.get(product.product_id).price
        return `<h5>${mapProduct.get(product.product_id).name}</h5>
        <p>Precio unitatio - ${mapProduct.get(product.product_id).price.toFixed(2)}</p>
        <p>Cantidad - ${product.amount}</p>
        <p>Precio total - ${(product.amount*mapProduct.get(product.product_id).price).toFixed(2)}</p>`
    }else{
        total += product.amount*mapProductUpdate.get(product.product_id).price
        return `<h5>${mapProductUpdate.get(product.product_id).name}</h5>
        <p>Precio unitatio - ${mapProductUpdate.get(product.product_id).price.toFixed(2)}</p>
        <p>Cantidad - ${product.amount}</p>
        <p>Precio total - ${(product.amount*mapProductUpdate.get(product.product_id).price).toFixed(2)}</p>`
    }
}

function completeOrder(orderID, payID) {
    fetch(`http://localhost:80/api/v1/order/${orderID}`, {
        body:JSON.stringify({"pay_id":payID}),
        method:"PATCH",
        headers:{'Content-Type': 'application/json'},
    }).then(res => {
        if (res.ok){
            deleteElementByID(`order-pending-${orderID}`);
            mapTableOrder.delete(mapOrderTable.get(orderID));
            mapOrderTable.delete(orderID);
            showSuccessful();
        }else{
            res.obj().then(x => showError(translateError(x.message)));
        }})
}

function completeOrderRemote(orderID) {
    fetch(`http://localhost:80/api/v1/order/remote/${orderID}`, {
        method:"PATCH",
        headers:{'Content-Type': 'application/json'},
    }).then(res => {
        if (res.ok){
            deleteElementByID(`order-pending-e-${orderID}`);
            mapTableOrder.delete(mapOrderTable.get(orderID));
            mapOrderTable.delete(orderID);
            showSuccessful();
        }else{
            res.obj().then(x => showError(translateError(x.message)));
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
    if(op.order.products != null){
        op.order.products.forEach(val => {
        if(!mapProductUpdate.has(val.id)){
            mapProductUpdate.set(val.id, val)
        }
        })
    }
        
    op.order_products.forEach(val => {orderProduct += loadOrderProduct(val)});
    
    order = `<div class="col-lg-3" id="order-pending-${op.order.id}">
    <div class="box">
    <div id="box-pending-${tableID}">
        <input id="i-p-${tableID}" hidden value="${total}"></input>
        <span>Mesa: ${tableID}</span>
        <h4>Encabezado</h4>
        <p>Fecha: ${timeToString(op.order.created_at)}</p>
        <p id="p-${tableID}">Precio Total: ${total.toFixed(2)}</p>
        <h4>Productos</h4>`
    order += orderProduct
    order += `</div><div class="text-center"><button type="button" onclick="openModalParamInput('completeOrder','Pedido',${op.order.id})">Completar pedido</button></div></div></div>`;
    return order;
}

function loadOrderSection(o) {
    let temp = 
    `<div class="container">
    <div class="section-title">
        <h2>Mis Pedidos <span>Pendientes</span></h2>
        <p>Aquí podras ver todos los pedidos sin completar que has realizado</p>
    </div>
    <div class="row" id="row-order">`;
    if (o != null){
        o.forEach((op,i) => {temp += loadOrder(op, i)});
    }
    temp += `</div></div>`;
    document.getElementById("my-incomplete-orders").innerHTML=temp
    document.getElementById("nav-my-order-pending").innerHTML = `<a href="#my-incomplete-orders">Pedidos Pendientes</a>`
}


function updateTotal(amount, price, id, item, key) {
    if (amount > 0){
        document.getElementById(id).innerText= `Precio Total: ${(amount*price).toFixed(2)}`;
        myMap.set(key, amount);
    }else{
        let node = document.getElementById(item);
        node.parentNode.removeChild(node);
        myMap.delete(key);
    }
}

function showSelectEstablishments() {
    let temp = `<option value="" selected disabled hidden>Selecciona un establecimiento</option>`;
    let temp2 = ``;
    temp += `<option value="0">Sin establecimiento</option>`;
    establishments.forEach(a => {
        temp += `<option value="${a.id}" >${a.id} / ${a.address.city}/ ${a.address.line1}/ ${a.address.postal_code}</option>`;
        temp2 += `<option value="${a.id}" >${a.id} / ${a.address.city}/ ${a.address.line1}/ ${a.address.postal_code}</option>`;
    });
    document.getElementById("hire-select-establishment").innerHTML = temp;
    temp2 += `<option value="0" selected >Nuevo Establecimiento</option>`
    document.getElementById("crud-select-establishment").innerHTML = temp2;
}

/*function SetEstablishments(a) {
    establishments = a;
    showHireUser()
    showFireUser()
}*/

function loadAllEstablishments(){
    fetch(`http://localhost:80/api/v1/establishment/`,{method:"GET"}).then(res => res.json().then(data => {
        establishments = data;
        data.forEach(val => mapEstablishments.set(val.id,val));
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
    let body;
    if (document.getElementById("hire-select-rol").value == "" ||document.getElementById("hire-select-establishment").value == "" || email=="" ){
        return `Necesitas llenar el formulario`;
    }
    if (document.getElementById("hire-select-establishment").value == "0"){
        if(rol_id > 2){return `Este role necesita de un establecimiento`};
        body = {"rol_id":rol_id};
    }else{
        body = {
            "rol_id":rol_id,
            "establishment_id":establishment_id,
        }
    }
    fetch(`http://localhost:80/api/v1/admin/hire/${email}`,{
        method:"PATCH",
        headers:{
            'Content-Type': 'application/json'
        },
        body:JSON.stringify(body)
    }).then(res => {
        if (!res.ok){
            res.json().then(x => err = window.alert(translateError(x.message)))
        }else{
            document.getElementsByName("hire-form").forEach(x => x.value="");
            showSuccessful();
        }
    });
}

function hireUserInStablishment(){
    let rol_id = parseInt(document.getElementById("hire-select-rol-st").value);
    let email = document.getElementById("hire-input-email-st").value;
    if(email == "" || document.getElementById("hire-select-rol-st").value == ""){return "Necesitas llenar el formulario";}
    fetch(`http://localhost:80/api/v1/user/hire/${email}`,{
        method:"PATCH",
        headers:{
            'Content-Type': 'application/json'
        },
        body:JSON.stringify({
            "rol_id":rol_id
        })
    }).then(res => showResultFunction(res));
}


function crudEstablishment(){
    let address = document.getElementsByName("form-crud-establishment")
    let err=``;
    let amountTables = address.item(6).value;
    let st = address.item(7).value;
    if (amountTables < 1){return `Agrega al menos una mesa`}
    address.forEach(val => {if(val.value == ""){err = "Necesitas completar el formulario"}});
    if (err != ""){return err}
    let dataAddress = {
        "line1":address.item(0).value,
        "line2":address.item(1).value,
        "city":address.item(2).value,
        "state":address.item(3).value,
        "country":address.item(4).value,
        "postal_code":address.item(5).value
    };
    let myInit = {method:"POST",headers:{'Content-Type': 'application/json'}};
    let addressInit={
        method:"",
        body:JSON.stringify(dataAddress),
        headers:{'Content-Type': 'application/json'}
    };
    if (st == "0"){
        addressInit.method = "POST";
        fetch("http://localhost:80/api/v1/address/", addressInit).then(res => res.json().then(data =>{
            myInit.body = JSON.stringify({address_id:data});
            fetch(`http://localhost:80/api/v1/establishment/${amountTables}`, myInit).then(res => {
                if(res.ok){showSuccessful()}else{res.json(x=>showError(translateError(x.message)))}
            });
        }))
    }else{
        addressInit.method = "PUT";
        fetch(`http://localhost:80/api/v1/establishment/${st}`, addressInit).then(res => showResultFunction(res));
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
                    <input class="form-control" id="${idInput}" type="number" min=0 onchange="updateTotal(this.value, ${price.toFixed(2)}, '${id}', '${idItem}', ${key})">
                </div>
                <div align="center">
                    <label id="${id}">Precio Total: ${price.toFixed(2)}</label>
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

function addAdminProduct(key, amount, src, name, price, description) {
    document.getElementById("input-name-u").value = name;
    document.getElementById("input-price-u").value = price;
    document.getElementById("img-u").src = src;
    document.getElementById("input-description-u").value = description;
    document.getElementById("hidden-u").value = key
}

function loadMenuProduct(data, funcName) {
    var temp = 
    `<div class="col-lg-6 menu-item">
    <div class="menu-content">
    <button name="button-menu" onclick="${funcName}(${data.id}, 1, '${data.img}', '${data.name}', ${data.price}, '${data.description}')">${data.name}</button><span>$${data.price.toFixed(2)}</span>
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
    if(email == ""){
        return "Necesitas llenar el formulario";
    }
    fetch(`http://localhost:80/api/v1/admin/fire/${email}`,{
        method:"PATCH",
    }).then(res => showResultFunction(res));
}

function fireUserInStablishment(){
    let email = document.getElementById("fire-input-email-st").value;
    if(email == ""){
        return "Necesitas llenar el formulario";
    }
    fetch(`http://localhost:80/api/v1/user/fire/${email}`,{
        method:"PATCH",
    }).then(res => showResultFunction(res));
}

function menuEmployee(data){
    let temp = "";
    data.forEach(item => temp+=loadMenuProduct(item, 'addProduct')); 
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
    data.forEach(item => temp+=loadMenuProduct(item, 'addAdminProduct')); 
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
        if(myAdmin.rol_id <= 2){
            menuAdmin(data)
        }else{
            menuEmployee(data)
        }
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
    if(products.length < 1){
        return `Selecciona al menos un producto`;
    }
    if(document.getElementById("select-table").value == ""){return `Selecciona una mesa`;}
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
                document.getElementById("shopping-cart").innerHTML = "";
                myMap.clear();
                document.getElementById("select-table").value = "";
                if (res.ok){
                    res.json().then(data => addOrderToPending(data));
                    showSuccessful();
                }else{
                    res.json().then(data => showError(translateError(data.message)));
                }
                })
                
                // Limpiar formulario, agregar a ordenes pendientes
    }
    }

function showMyOrder(){
    document.getElementById("nav-my-order").innerHTML = `<a href="#my-order">Mi pedido</a>`
    document.getElementById("my-order").innerHTML = `<div class="container-fluid">
    <div class="section-title">
        <h2>Nuevo <span>Pedido</span></h2>
        <p>Aquí se listaran todos los productos seleccionados para proceder con su respectivo pedido</p>
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
    document.getElementById("nav-create-establishment").innerHTML = `<a href="#create-establishment">Establecimientos</a>`
    document.getElementById("create-establishment").innerHTML = `<div class="container">
    <div class="section-title">
    <h2>Registrar/Actualizar <span>Establecimiento</span></h2>
    <p>Agrega/Actualiza un establecimiento</p>
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
    <div class="text-center" onclick="openModal('crudEstablishment', 'Establecimiento')"><button type="button">Enviar establecimiento</button></div>
    </form>
    </div>`
}



function showFireUser(){
    document.getElementById("nav-fire").innerHTML = `<a href="#fire-employee">Despedir</a>`
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
    document.getElementById("nav-fire-st").innerHTML = `<a href="#fire-employee-st">Despedir</a>`
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
    document.getElementById("nav-hire-st").innerHTML = `<a href="#hire-employee-st">Contratar</a>`
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
    document.getElementById("nav-hire").innerHTML = `<a href="#hire-employee">Contratar</a>`
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

function deleteProduct() {
    let id = parseInt(document.getElementById("hidden-u").value);
    let err = ``;
    if (document.getElementById("hidden-u").value == "" || id < 1){
        return `Selecciona un producto a eliminar`;
    }
    fetch(`http://localhost:80/api/v1/product/${id}`,{
        method: "DELETE"
    }).then(res => {
        showResultFunction(res);
        $('#form-set-image').trigger("reset");
        document.getElementById("img-u").src =``;
        document.getElementById("input-name-u").value = "";
        document.getElementById("input-price-u").value = "";
        document.getElementById("input-description-u").value="";
        document.getElementById("hidden-u").value="";
        document.getElementById("input-file-u").value="";
        });
}

function updateProduct() {
    let name = document.getElementById("input-name-u").value;
    let price = parseFloat(document.getElementById("input-price-u").value);
    let description = document.getElementById("input-description-u").value;
    var file = document.getElementById("input-file-u").files[0];
    let id = parseInt(document.getElementById("hidden-u").value);
    console.log(id == NaN);
    if (document.getElementById("hidden-u").value == "" || id < 1){return `Selecciona un producto a actualizar`;}
    if (name == "" || document.getElementById("input-price-u").value == "" || description == "" || file == undefined){
        return `Necesitas llenar el formulario`;
    }
    fetch(`http://localhost:80/api/v1/product/${id}`,{
        method: "PUT",
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
            sendFormSetImageU();
        }
        showResultFunction(res);
        $('#form-set-image').trigger("reset");
        document.getElementById("img-u").src =``;
        document.getElementById("input-name-u").value = "";
        document.getElementById("input-price-u").value = "";
        document.getElementById("input-description-u").value="";
        document.getElementById("hidden-u").value="";
        document.getElementById("input-file-u").value="";
        });
}

function createProduct() {
    let name = document.getElementById("input-name").value;
    let price = parseFloat(document.getElementById("input-price").value);
    let description = document.getElementById("input-description").value;
    var file = document.getElementById("input-file").files[0];
    if (name == "" || document.getElementById("input-price").value == "" || description == "" || file == undefined){
        return `Necesitas completar el formulario`;
    }
    if(price <= 0){return `El precio tiene que ser mayor a cero`};
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
        }
        showResultFunction(res)
        $('#form-set-image').trigger("reset");
        document.getElementById("img").src =``;
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
}

function sendFormSetImage(){
    nameImage = document.getElementById("input-name").value;
    document.getElementById("form-set-image").action = `http://localhost:80/api/v1/product/img/${nameImage}`;
    document.getElementById("form-set-image").submit();
}

function updateNameImageU(){
    nameUpdateImage = document.getElementById("input-name-u").value
    document.getElementById("submit-product-u").formAction = `http://localhost:80/api/v1/product/img/${nameUpdateImage}`
}

function sendFormSetImageU(){
    nameUpdateImage = document.getElementById("input-name-u").value;
    document.getElementById("form-set-image-u").action = `http://localhost:80/api/v1/product/img/${nameUpdateImage}`;
    document.getElementById("form-set-image-u").submit();
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

$('#form-set-image-u').submit(function () {
    $.ajax({
        type: $('#form-set-image-u').attr('method'), 
        url: `"http://localhost:80/api/v1/product/img/${nameUpdateImage}"`,
        data: $('#form-set-image-u').serialize(),
        success: function (data) { console.log("Enviado") }  
    });
    return false;
});

function showUpdateProduct(){
    document.getElementById("nav-update-product").innerHTML = `<a href="#update-product">Actualizar Producto</a>`
    document.getElementById("update-product").innerHTML=`<div class="container">
        <div class="section-title">
        <h2>Actualizar <span>Productos</span></h2>
        <p>Actualización/Eliminación de productos</p>
    </div>
    </div>
    <div class="container book-a-table">
    <form class="php-email-form" method="post" enctype="multipart/form-data" id="form-set-image-u" target="request">
        <div class="form-row" >
        <div class="col-md-6 form-group" id="div-input-name">
            <input type="text" name="name" class="form-control" id="input-name-u" oninput="updateNameImageU()" placeholder="Nombre del producto">
        </div>
        <div class="col-md-6 form-group">
            <input type="number" class="form-control" name="price" id="input-price-u" placeholder="Precio" >
        </div>
        <div class="col-md-6 form-group">
            <input class="form-control" name="file" type="file" id="input-file-u" accept="image/*" onchange="mostrarU()" >
        </div>
        <input type="hidden" id="hidden-u">
        <div class="gallery-item">
            <img id="img-u" class="img-fluid">
        </div>
        </div>
        <div class="form-group">
        <textarea id="input-description-u" class="form-control" name="message" rows="5" data-rule="required" data-msg="Please write something for us" placeholder="Descripción"></textarea>
        </div>
        <div class="text-center">
            <button type="button" onclick="openModal('updateProduct', 'Actualizar Producto')">Actualizar</button>
            <button type="button" onclick="openModal('deleteProduct', 'Eliminar Producto')">Eliminar</button>
        </div>
        
    </form>
    </div>
    
    <iframe name="request" width="0" height="0" frameborder="0" id="request-u" style="display: none;"></iframe>`
    
}


function showCreateProduct(){
    document.getElementById("nav-new-product").innerHTML=`<a href="#create-product">Crear Producto</a>`
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
    showUpdateProduct();
    getAllRoles(2);
    fetchPayments();
    //showMyOrder();
}

function showMakerOrderRemote(){

}

function showFunctionByPermission(id){
    console.log(id)
    switch (id) {
        case 2:
            showCRUDEstablishment();
            break
        case 3:
            getAllRoles(2);
            break;
        case 4:
            showCreateProduct();
            showUpdateProduct();
            break;
        case 5:
            loadMyEstablishments();
            break;
        case 6:
            break;
        case 7:
            getAllOrdersPendingByEstablishment();
            break;
        case 8:
            fetchPayments();
            break;
        case 9:
            getAllOrdersByEstablishment();
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
    loadPayment();
    modal = document.getElementById("myModal");
    isLogin();
    showNavResponse();
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

function mostrarU(){
    var file = document.getElementById("input-file-u").files[0];
    var reader = new FileReader();
    if (file) {
        reader.readAsDataURL(file );
        reader.onloadend = function () {
            document.getElementById("img-u").src = reader.result;
        }
        }
}