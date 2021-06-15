let myMap = new Map();
let mapProduct = new Map();
let mapProductUpdate = new Map();
let myOrders;
let myUser;
let total=0;
let last=0;


function fetchEstablishment() {
    fetch("http://localhost:80/api/v1/establishment/").then(res => res.json().then(obj => {
        let temp = ``;
        obj.forEach(val => temp += `<option value="${val.id}">${val.address.line1} / ${val.address.city} / ${val.address.postal_code}</option>`);
        document.getElementById("select-establishment").innerHTML += temp;
    }))
}

function updateUserEmailAndPassword() {
    let pwd = document.getElementById("password-change-input").value
    let data = {
        "password":pwd
    }
    fetch("http://localhost:80/api/v1/user/password/", {
        //credentials: 'same-origin',
        method: 'PATCH',
        body: JSON.stringify(data),
        headers:{
            'Content-Type': 'application/json'
        }}).then(response => console.log(response))
}

function loadOrderProduct(product) {
    if(mapProduct.has(product.product_id)){
        total += product.amount*mapProduct.get(product.product_id).price
        return `<h5>${mapProduct.get(product.product_id).name}</h5>
        <p>Precio unitatio - ${mapProduct.get(product.product_id).price}</p>
        <p>Cantidad - ${product.amount}</p>
        <p>Precio total - ${product.amount*mapProduct.get(product.product_id).price}</p>`
    }else{
        total += product.amount*mapProductUpdate.get(product.product_id).price
        return `<h5>${mapProductUpdate.get(product.product_id).name}</h5>
        <p>Precio unitatio - ${mapProductUpdate.get(product.product_id).price}</p>
        <p>Cantidad - ${product.amount}</p>
        <p>Precio total - ${product.amount*mapProductUpdate.get(product.product_id).price}</p>`
    }
    
}

function loadOrder(op, i) {
    console.log(op);
    let order; 
    let orderProduct = ``
    total=0;
    if(!mapOrders.has(op.order.id)){
        mapOrders.set(op.order.id, op);
    }
    if(op.order.products != null){
        op.order.products.forEach(val => {
            if(!mapProductUpdate.has(val.id)){
                mapProductUpdate.set(val.id, val)
            }
        })
    }
    
    op.order_products.forEach(val => {orderProduct += loadOrderProduct(val)});
    order = `<div class="col-lg-3" onclick="printPDF(${op.order.id})">
    <div class="box">
        <span>${i + 1}</span>
        <h4>Encabezado</h4>
        <p>Fecha: ${timeToString(op.order.created_at)}</p>
        <p>Tipo de pago: ${mapPayments.get(op.order.pay_id)}</p>
        <p>Precio Total: ${total}</p>
        <h4>Productos</h4>`
    order += orderProduct
    order += `</div></div>`;
    last += 1;
    return order;
}

function loadOrderSection(o) {
    let temp = 
    `<div class="container">

    <div class="section-title">
        <h2>Mis <span>Pedidos</span></h2>
        <p>Aqui podras ver todos los pedidos que has realizado</p>
    </div>
    <div class="row" id="row-order">`;
    if(o != null) {
        o.forEach((op,i) => {temp += loadOrder(op, i)});
    }
    temp += `</div></div>`;
    document.getElementById("orders").innerHTML=temp;
}

function loadAllOrders() {
    fetch("http://localhost:80/api/v1/order/", {
        method: 'GET',
        headers:{
            'Content-Type': 'application/json'
        }
    }).then(res => {
        if (res.status != 204 && res.ok){
            res.json().then(data => {
            setMyOrders(data).then(loadPay(data));})
        }else{
            setMyOrders(null).then(loadPay(null));
        }
    }
        
    ).catch(a => console.log(a));
}
async function setMyOrders(data){
    myOrders = data;
}

function loadPay(data){
    if(mapPayments.size == 0){
        fetch("http://localhost:80/api/v1/pay/").then(res => {
            res.json().then(data => {
                data.forEach(x => mapPayments.set(x.id, x.name));
                loadOrderSection(data);
            })}
        );
    }else{
        loadOrderSection(data);
    }
}

async function loadPaymentMethod(){
    var myInit = {method: 'GET'};
    var myRequest = new Request("http://localhost:80/api/v1/pay/", myInit);
    fetch(myRequest).then(res => {
        res.json().then(
        data => {
            if (data.length > 0) {
                let temp=`<option value="" selected disabled hidden>Selecciona un metodo de pago</option>`;
                data.forEach((itemData) => {
                    temp += `<option value=${itemData.id}>${itemData.name}</option>`
                    mapPayments.set(itemData.id, itemData.name)
                });
                document.getElementById("select-payment").innerHTML = temp;
                myPayments = data
            }
        })
    })
}

function makeOrder(){
    let address = document.getElementsByName("form-my-order")
    let isValid = true;
    if(myMap.size == 0){return `Necesitas agregar al menos un producto al pedido`}
    address.forEach((value) => {
        if (value.value == ""){
            isValid = false;
        }
    })
    if (!isValid){return `Necesitas completar el formulario`}
    let dataAddress = {
        "line1":address.item(0).value,
        "line2":address.item(1).value,
        "city":address.item(2).value,
        "state":address.item(3).value,
        "country":address.item(4).value,
        "postal_code":address.item(5).value
    };
    let idAddress;
    let products = [];
    fetch("http://localhost:80/api/v1/address/", {
        //credentials: 'same-origin',
        method: 'POST',
        body: JSON.stringify(dataAddress),
        headers:{
            'Content-Type': 'application/json'
        }}).then(response => {
            if (response.ok){
                response.json().then(data => {
                idAddress = data;
                myMap.forEach((val, key) => {
                    let p = {
                        "product_id":parseInt(key),
                        "amount":parseInt(val)
                    };
                    products.push(p);
                });
                let orderProduct = {
                    "order":{
                        "pay_id":parseInt(address.item(6).value),
                        "establishment_id":parseInt(address.item(7).value),
                        "address_id":idAddress
                    },
                    "order_products":products
                };
                fetch("http://localhost:80/api/v1/order/remote/", {
                    //credentials: 'same-origin',
                    method: 'POST',
                    body: JSON.stringify(orderProduct),
                    headers:{
                        'Content-Type': 'application/json'
                    }}).then(res => {
                        if (res.ok){
                            document.getElementById("shopping-cart").innerHTML = "";
                            myMap.clear();
                            address.forEach((v) => v.value = ``);
                            res.json().then(obj => document.getElementById("row-order").innerHTML += loadOrder(obj, last));
                        }
                        showResultFunction(res);
                    });
        })}else{
            response.json().then(data => showError(translateError(data.message)));
        }
    });
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

// add product to order
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

function getMenu(url, fn){
    var myInit = {method: 'GET'};
    //headers: myHeaders
    var myRequest = new Request(url, myInit);
    //var myHeaders = myRequest.headers;
    fetch(myRequest).then(res => {
        res.json().then(
        data => {
            if (data.length > 0) {
                var temp = "";
                data.forEach((itemData) => {
                    temp += fn(itemData)
                });
                document.getElementById("menu").innerHTML=`<div class="container">
                <div class="section-title">
                    <h2>Revisa nuestro sabroso <span>Menú</span></h2>
                </div>
                <div class="row menu-container" id="menu-container">
                ${temp}
                </div>
        
                </div>`;
            }
        })
    })
}
function LogOut() {
    document.cookie = "" 
    sessionStorage.clear()
    fetch("http://localhost/api/v1/user/logout/",{method:"GET"}).then(location.reload())
}

function post(url) {
    var email = document.getElementById('email').value;
    var password = document.getElementById('password').value;
    //http.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
    let data = {
        "email":email,
        "password":password
    }
    fetch(url, {
        //credentials: 'same-origin',
        method: 'POST', // or 'PUT'
        body: JSON.stringify(data), // data can be `string` or {object}!
        headers:{
            'Content-Type': 'application/json'
        }
        }).then(response => {
            if (response.ok){
                document.cookie = response.headers.get("Cookie");
                response.json().then(data => {sessionStorage.setItem("authorization",data.token);});
                location.reload();
            }else{
                showError("Credenciales no validas")
            }
            
        }).catch(err => console.log(err));

}

function caseNotLogin() {
    let temp = `<div class="container">
    <div class="section-title" action="" method="POST">
        <h2><span>Registrar/Logear usuario</span></h2>
        <p>Por favor llena el siguiente formulario y da click en uno de los botones para crear una cuenta nueva o acceder a tu cuenta existente</p>
    </div>
    <form class="php-email-form" action="">
        <div class="form-row">
        <div class="col-lg-4 col-md-6 form-group">
            <input type="email" class="form-control" name="email" id="email" placeholder="Tu Correo Electrónico" oninput="return validar_email(email.value)" data-rule="email" data-msg="Please enter a valid email">
            <span id="email-error">Email no valido</span>
            <div class="validate"></div>
        </div>
        <div class="col-lg-4 col-md-6 form-group">
            <input type="password" class="form-control" name="password" id="password" placeholder="Ingresa una contraseña" oninput="return validar_clave(password.value)" data-msg="Please enter at least 4 chars">
            <span id="password-error">Contraseña no valida</span>
            <div class="validate"></div>
        </div>
        <div class="text-center"><button  type="button" disabled id="button-singup" onclick="post('http://localhost:80/api/v1/user/')">Registrar</button></div>
        <div class="text-center"><button type="button" disabled id="button-login" onclick="post('http://localhost:80/api/v1/user/login/')">LogIn</button></div>
    </form>
    </div>`
    document.getElementById("book-a-table").innerHTML = temp;
    let x = document.getElementsByName("button-menu");
    x.forEach((a, b,c ) => a.disabled=true)
}

function sectionMyOrder() {
    fetch("http://localhost:80/api/v1/establishment/").then(res => res.json().then(obj => {
        let temp = ``;
        obj.forEach(val => temp += `<option value="${val.id}">${val.address.line1} / ${val.address.city} / ${val.address.postal_code}</option>`);
        document.getElementById("gallery").innerHTML = `<div class="container-fluid">
        <div class="section-title">
            <h2>Mi <span>Pedido</span></h2>
            <p>Aqui se listaran todos los productos seleccionados para proceder con su respectivo pedido</p>
        </div>
        <div class="container">
            <div id="shopping-cart" class="row no-gutters"></div>
            <form class="php-email-form" action="">
            <div class="form-row">
                <div class="col-lg-4 col-md-6 form-group">
                <input name="form-my-order" type="text" class="form-control" id="line1" placeholder="Ingresa tu Calle y numero de casa" required>
                </div>
                <div class="col-lg-4 col-md-6 form-group">
                <input name="form-my-order" type="text" class="form-control" id="line2" placeholder="Ingresa tu Colonia">
                </div>
                <div class="col-lg-4 col-md-6 form-group">
                <input name="form-my-order" type="text" class="form-control" id="city" placeholder="Ingresa tu Ciudad" required>
                </div>
                <div class="col-lg-4 col-md-6 form-group">
                <input name="form-my-order" type="text" class="form-control" id="state" placeholder="Ingresa tu Estado" required>
                </div>
                <div class="col-lg-4 col-md-6 form-group">
                <input name="form-my-order" type="text" class="form-control" id="country" placeholder="Ingresa tu Pais" value="Mexico" required>
                </div>
                <div class="col-lg-4 col-md-6 form-group">
                <input name="form-my-order" type="text" class="form-control" id="postal_code" placeholder="Codigo Postal" required>
                </div>
                <div class="col-lg-4 col-md-6 form-group">
                <div class="text-center">
                <select name="form-my-order" id="select-payment" required>
                </select>
                </div>
                </div>
                <div class="col-lg-4 col-md-6 form-group">
                <div class="text-center">
                <select name="form-my-order" id="select-establishment" required>
                <option value="" selected disabled hidden>Selecciona un Establecimiento</option>
                ${temp}
                </select>
                </div>
                </div>
                <div class="col-lg-4 col-md-6 form-group">
                <div class="text-center"><button type="button" onclick="openModal('makeOrder', 'Pedido')">Realizar pedido</button></div></div>
            </div>
            </form>
        </div>`;
        loadPaymentMethod().then(() => loadAllOrders());
    }))
}

function htmlSectionAccount(){
    document.getElementById("nav-account").innerHTML = `<a href="#account">Mi Usuario</a>`
    document.getElementById("account").innerHTML = `<div class="container">
    <div class="section-title" action="" method="POST">
        <h2><span>Cuenta de Usuario</span></h2>
        <p>Aqui podras cambiar tu contraseña y desconectar tu usuario actual</p>
    </div>
    <form class="php-email-form" action="">
        <div class="form-row">
        <div class="col-lg-4 col-md-6 form-group">
            <input type="password" class="form-control" id="password-change-input" name="password"  placeholder="Ingresa una contraseña" oninput="isChangePasswordValid()">
            <span id="password-change-error">Contraseña no valida</span>
        </div>
        <div class="text-center"><button id="password-change" type="button" disabled onclick="updateUserEmailAndPassword()">Cambiar contraseña</button></div>
    </form>
    </div>
    <div class="container">
    <form class="php-email-form" action="">
    <div class="text-center"><button id="button-logout" type="button" onclick="LogOut()">Cerrar sesion</button></div>
    </form>
    </div>`
}

function caseLogin() {
    
    sectionMyOrder();
    document.getElementById("li-my-order").innerHTML = `<a href="#gallery">Mi Pedido</a>`;
    document.getElementById("nav-orders").innerHTML = `<a href="#orders">Pedidos Realizados</a>`;
    htmlSectionAccount();
}

async function setMyUser(data){
    myUser = data
    if (myUser.rol_id != null) {
        document.getElementById("nav-admin").innerHTML = `<a href="admin.html">Admin</a>`
    }
} 

function switchCaseSession() {
    var myInit = {method: 'GET'};
    var myRequest = new Request('http://localhost:80/api/v1/user/login/', myInit);
    fetch(myRequest).then(res => {
        if (res.ok){
            var temp = `<a href="#menu">Realizar pedido</a>`;
            res.json().then(a =>{
                setMyUser(a).then(console.log(myUser));
                caseLogin();
            })
        }else{
            var temp = `<a href="#book-a-table">Registrarse</a>`
            caseNotLogin()
        }
        document.getElementById("nav-item-session").innerHTML = temp;
    })
}

function validar_email(email){
    let re = /^[^\s@]+@[^\s@]+\.[^\s@]+$/
    var a = re.test(email);
    if (!a){
        document.getElementById("email-error").innerHTML = "Email no valido";
    }
    else{
        document.getElementById("email-error").innerHTML = "";
    }
    isValid()
}

function isPassword(password) {
    if(password.length >= 8)
    {		
        var mayuscula = false;
        var minuscula = false;
        var numero = false;
        var caracter_raro = false;
        
        for(var i = 0;i<password.length;i++)
        {
            if (mayuscula && minuscula && numero && caracter_raro) {
                return true;
            }
            if(password.charCodeAt(i) >= 65 && password.charCodeAt(i) <= 90)
            {
                mayuscula = true;
            }
            else if(password.charCodeAt(i) >= 97 && password.charCodeAt(i) <= 122)
            {
                minuscula = true;
            }
            else if(password.charCodeAt(i) >= 48 && password.charCodeAt(i) <= 57)
            {
                numero = true;
            }
            else
            {
                caracter_raro = true;
            }
        }
        return mayuscula && minuscula && numero && caracter_raro;
    }
    return false;
}

function isChangePasswordValid(){
    password = document.getElementById("password-change-input").value
    if (isPassword(password)){
        console.log(password)
        document.getElementById("password-change").disabled=false
        document.getElementById("password-change-error").innerHTML = "";
    }else{
        document.getElementById("password-change").disabled=true
        document.getElementById("password-change-error").innerHTML = "Contraseña no valida";
    }
}

function validar_clave(contrasenna){
			if(contrasenna.length >= 8)
			{		
				var mayuscula = false;
				var minuscula = false;
				var numero = false;
				var caracter_raro = false;
				
				for(var i = 0;i<contrasenna.length;i++)
				{
                    if (mayuscula && minuscula && numero && caracter_raro) {
                        document.getElementById("password-error").innerHTML = "";
                    }
					if(contrasenna.charCodeAt(i) >= 65 && contrasenna.charCodeAt(i) <= 90)
					{
						mayuscula = true;
					}
					else if(contrasenna.charCodeAt(i) >= 97 && contrasenna.charCodeAt(i) <= 122)
					{
						minuscula = true;
					}
					else if(contrasenna.charCodeAt(i) >= 48 && contrasenna.charCodeAt(i) <= 57)
					{
						numero = true;
					}
					else
					{
						caracter_raro = true;
					}
				}
                a = mayuscula && minuscula && numero && caracter_raro
                if (!a){
                    document.getElementById("password-error").innerHTML = "Contraseña no valida";
                }else{
                    document.getElementById("password-error").innerHTML = "";
                }
			}else{
                document.getElementById("password-error").innerHTML = "Contraseña no valida";
            }
            isValid()
}

function isValid() {
    if (document.getElementById("password-error").innerHTML === "" && document.getElementById("email-error").innerHTML === "") {
        document.getElementById("button-singup").disabled = false;
        document.getElementById("button-login").disabled = false;
    }else{
        document.getElementById("button-singup").disabled = true;
        document.getElementById("button-login").disabled = true;
}
}
    
function getAll() {
    modal = document.getElementById("myModal");
    getMenu('http://localhost:80/api/v1/product/', loadMenuProduct);
    switchCaseSession();
}