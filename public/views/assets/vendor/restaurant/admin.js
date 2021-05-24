
let myAdmin;
let nameImage;
let myRole;
let roles;
let establishments;

const PERMISSIONS = {
    WhitOutRestriction:1,

}
function showSelectEstablishments() {
    let temp = `<option value="" selected disabled hidden>Selecciona un establecimiento</option>`;
    establishments.forEach(a => {
        temp += `<option value="${a.id}" >${a.id}</option>`;
    });
    document.getElementById("hire-select-establishment").innerHTML = temp;
}

/*function SetEstablishments(a) {
    establishments = a;
    showHireUser()
    showFireUser()
}*/

function loadAllEstablishments(){
    fetch(`http://localhost:80/api/v1/establishment/`,{method:"GET"}).then(res => res.json().then(data => {
        console.log("rol: ",data);
        establishments = data;
        showHireFireAdmin();
    })).catch(a => console.log(a));
}
function SetMyAdmin(a) {
    myAdmin = a;
    getMyRolePermissions();
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

function showHireFireAdmin(){
    showHireUser();
    showFireUser();
}
function showHireFireUser(){
    showHireUserInStablishment();
    showFireUserInStablishment();
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
        <div class="text-center"><button type="button" onclick="fireUser()">Despedir Empleado</button></div>
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
        <div class="text-center"><button type="button" onclick="fireUserInStablishment()">Despedir Empleado</button></div>
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
        <div class="text-center"><button type="button" onclick="hireUserInStablishment()">Contratar Empleado</button></div>
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
        <div class="text-center"><button type="button" onclick="hireUser()">Contratar Empleado</button></div>
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
        <input id="submit-product" type="button" value="Submit" onclick="sendFormSetImage()">
        <div class="text-center"><button type="button" onclick="createProduct()">Registrar producto</button></div>
    </form>
    </div>
    
    <iframe name="request" width="0" height="0" frameborder="0" id="request" style="display: none;"></iframe>`
    
}

function showAll(){
    showCreateProduct();
    getAllRoles(2);
}

function showMakerOrderRemote(){

}

function showFunctionByPermission(id){
    console.log(id)
    switch (id) {
        case 2:
            showCreateProduct();
            break
        case 3:
            getAllRoles(2);
            break;
        case 4:
            break;
        case 5:
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