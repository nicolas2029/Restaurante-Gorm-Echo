let myAdmin;

async function SetMyAdmin(a) {
    myAdmin = a;
}

function setImageToProduct(id, name) {
    let file = document.getElementById(id).file;
    fetch(`http://localhost:80/api/v1/product/img/${name}/`,{
        method:"POST",
        body:file,
        headers:{
            'Content-Type': `multipart/form-data`
        }
    }).then(console.log("IMAGEN AGREGADA CON EXITO"));

}

async function createProduct() {
    let name = document.getElementById("input-name").value;
    let price = documenrt.getElementById("input-price").value;
    fetch("http://localhost:80/api/v1/product",{
        method: "POST",
        body:stringify({
            "name":name,
            "price":price
        }),
        headers:{
            'Content-Type': 'application/json'
        }
    }).then(res => {
        if (res.ok){
            setImageToProduct("img", name);
        }
        else{
            console.log(res.status);
        }
    });
}

async function isLogin (){
    var myInit = {method: 'GET'};
    var myRequest = new Request('http://localhost:80/api/v1/user/login/', myInit);
    fetch(myRequest).then(res => {
        if (res.ok){
            res.json().then(a =>{
                SetMyAdmin(a).then(getRolPermissions());
            })
        }
    })
}

function showCreateProduct(){
    document.getElementById("create-product").innerHTML=`<div class="container">
        <div class="section-title">
        <h2>Registrar <span>Productos</span></h2>
        <p>Registro de productos nuevos</p>
    </div>
    </div>
    <div class="container book-a-table">
    <form class="php-email-form" action="">
        <div class="form-row">
        <div class="col-md-6 form-group">
            <input type="text" name="name" class="form-control" id="input-name" placeholder="Nombre del producto">
        </div>
        <div class="col-md-6 form-group">
            <input type="number" class="form-control" name="price" id="input-price" placeholder="Precio" >
        </div>
        <div class="col-md-6 form-group">
            <input class="form-control" type="file" id="input-file" accept="image/*" onchange="mostrar()" >
        </div>
        <div class="gallery-item">
            <img id="img" class="img-fluid">
        </div>
        </div>
        <div class="form-group">
        <textarea class="form-control" name="message" rows="5" data-rule="required" data-msg="Please write something for us" placeholder="Descripción"></textarea>
        </div>
        <div class="text-center"><button type="button">Registrar producto</button></div>
    </form>
    </div>`
}

function showAll(){
    showCreateProduct()
}

function showMakerOrderRemote(){

}

function showAllFunctionsByRol(rol){
    if(rol.id == 1){
        showAll()
    }
}

async function getRolPermissions() {
    var myInit = {method: 'GET'};
    var myRequest = new Request(`http://localhost:80/api/v1/rol/${myAdmin.rol_id}`, myInit);
    fetch(myRequest).then(res => res.json()).then(data => {
        showAllFunctionsByRol(data);
    })
}

function loadAdminPage() {
    isLogin()
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