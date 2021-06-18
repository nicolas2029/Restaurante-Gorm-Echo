let mapOrders = new Map();
let totalOrder = 0;
let mapPayments = new Map();
let modal;
let selectPayments="";
let amountNav = 0;
let establishments;
let mapEstablishments = new Map();

function showNavResponse(){
    document.getElementsByClassName("mobile-nav-toggle d-lg-none")[0].addEventListener("click", () => {
        if(amountNav == 0){
            var $mobile_nav = $('.nav-menu').clone().prop({
            class: 'mobile-nav d-lg-none'
            });
            $('body').append($mobile_nav);
            amountNav = 1;
        }
    })
}

function showResultFunction(res){
    if(res.ok){
        showSuccessful();
        return true;
    }else{
        res.json().then(x => showError(translateError(x.message)));
        return false;
    }
}

function showSuccessful(){
    window.alert("COMPLETADO CON EXITO");
}

function showError(err){
    window.alert(err)
}

function closeModal(){
    modal.innerHTML = "";
    modal.style.display = "none";
}

function closeAndExecute(func){
    closeModal();
    let res = func();
    if (res != null) {
        showError(res);
    }
    
}

function closeAndExecuteParam(func, param){
    closeModal();
    let res = func(param);
    if (res != null) {
        showError(res);
    }
    
}

function closeAndExecuteParamInput(func, param){
    let id = document.getElementById("modal-select").value;
    if(id == ""){
        showError("Selecciona un método de pago");
        return
    }
    closeModal();
    let res = func(param, parseInt(id));
    if (res != null) {
        showError(res);
    }
    
}

function openModalParamInput(funcName, action, param){

    modal.innerHTML = `<div class="modal-content">
        <header class="close" id="modal-close" onclick="closeModal()">&times;</header>
        <div class="section-title">
        <h1 style="font-weight: bold">Confirmar<span style="color: #ffb03b"> ${action}</span></h1>
        <select id="modal-select"><option value="" selected="" disabled="" hidden="">Selecciona un método de pago</option>${selectPayments}</select>
        <p>¿Estás seguro de continuar?</p>
        </div>
        <div>
        <button type="button" onclick="closeModal()" class="cancelbtn">Cancelar</button>
        <button type="button" onclick="closeAndExecuteParamInput(${funcName}, ${param})" class="deletebtn">Confirmar</button>
        </div>
    </div>`
    modal.style.display = "block";
}

function openModalParam(funcName, action, param){
    modal.innerHTML = `<div class="modal-content">
        <header class="close" id="modal-close" onclick="closeModal()">&times;</header>
        <div class="section-title">
        <h1 style="font-weight: bold">Confirmar<span style="color: #ffb03b"> ${action}</span></h1>
        <p>¿Estás seguro de continuar?</p>
        </div>
        <div>
        <button type="button" onclick="closeModal()" class="cancelbtn">Cancelar</button>
        <button type="button" onclick="closeAndExecuteParam(${funcName}, ${param})" class="deletebtn">Confirmar</button>
        </div>
    </div>`
    modal.style.display = "block";
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
        <button type="button" onclick="closeAndExecute(${funcName})" class="deletebtn">Confirmar</button>
        </div>
    </div>`
    modal.style.display = "block";
}



window.onclick = function(event) {
    if (event.target == modal) {
        modal.style.display = "none";
    }
}


function translateError(key) {
    switch (key) {
        case `user not logged in`:
            return `Necesitas iniciar sesion`;
        case `user not found`:
            return `Usuario no encontrado`;
        case `not found`:
            return `Informacion no encontrada`;
        case `you are not autorized`:
            return `No tienes permizo para realizar esta accion`;
        case `user whitout role`:
            return `No tienes el role necesario`;
        case `invalid password`:
            return `Contraseña invalida`;
        case `empty result`:
            return `Sin resultados`;
        case `invalid email`:
            return `Email invalido`;
        case `email already in use`:
            return `Email actualmente en uso`;
        case `invalid role`:
            return `Role no valido`;
        case `cannot get data`:
            return `No se pudo obtener la informacion`;
        case `table not available`:
            return `Mesa no disponible`;
        case `empty order`:
            return `Pedido sin productos`;
        case `order already completed`:
            return `El pedido ya ha sido completado`;
        case `empty address`:
            return `Direccion vacia`;
        case `product already updated`:
            return `Producto actualmente actualizado`;
        default:
            return key;
    }
}

function loadPayment(){
    fetch("http://localhost:80/api/v1/pay/").then(res => res.json().then(data => {
        data.forEach(obj => {mapPayments.set(obj.id,obj.name);
            selectPayments += `<option id="op-modal-${obj.id}" value="${obj.id}">${obj.name}</option>`;})
    }));
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

function createOrderProduct(product){
    let data;
    if(mapProduct.has(product.product_id)){
        data = mapProduct.get(product.product_id);
    }else{
        data = mapProductUpdate.get(product.product_id);
    }
    totalOrder += product.amount*data.price
    return `<tr>
        <td class="precio">${product.amount}</td>
        <td class="producto">${data.name}</td>
        <td class="precio">$${data.price.toFixed(2)}</td>
        <td class="precio">$${(data.price * product.amount).toFixed(2)}</td>
        </tr>`
    
}

function createInvoice(order){
    let temp = ``;
    let pay = ``;
    let st = ``;
    totalOrder = 0;
    console.log(order)
    order.order_products.forEach(element => {
        temp += createOrderProduct(element);
    });
    if(mapPayments.has(order.order.pay_id)){
        pay = `<br>Tipo de pago: ${mapPayments.get(order.order.pay_id)}`;
    }
    if(mapEstablishments.has(order.order.establishment_id)){
        let t = mapEstablishments.get(order.order.establishment_id);
        st = `<br>Direccion:  ${t.address.city} / ${t.address.state} / ${t.address.country} / ${t.address.line1} / ${t.address.postal_code}`;
    }
    return `<div class="ticket">
        <p class="centrado">Beer para creer
            <br>${timeToString(order.order.created_at)}
            ${pay}
            ${st}
        </p>
        <table class="centrado">
            <thead>
                <tr>
                    <th class="producto">CANTIDAD</th>
                    <th class="producto">PRODUCTO</th>
                    <th class="precio">PRECIO C/U</th>
                    <th class="precio">PRECIO TOTAL</th>
                </tr>
            </thead>
            <tbody>
                ${temp}
                <tr>
                    <td class="precio"></td>
                    <td class="producto">TOTAL</td>
                    <td class="precio">$${totalOrder.toFixed(2)}</td>
                    <td class="precio"></td>
                </tr>
            </tbody>
        </table>
        <p class="centrado">¡GRACIAS POR SU COMPRA!
            <br>beerparacreer.com</p>
    </div>`
}


function printPDF(id){
    let order = mapOrders.get(id);
    console.log("order",order);

    const $elementoParaConvertir = createInvoice(order); // <-- Aquí puedes elegir cualquier elemento del DOM
    html2pdf()
        .set({
            margin: 1,
            filename: 'Ticket.pdf',
            image: {
                type: 'jpeg',
                quality: 0.98
            },
            html2canvas: {
                scale: 3, // A mayor escala, mejores gráficos, pero más peso
                letterRendering: true,
            },
            jsPDF: {
                unit: "in",
                format: "a3",
                orientation: 'portrait' // landscape o portrait
            }
        })
        .from($elementoParaConvertir)
        .save()
        .catch(err => console.log(err));
}