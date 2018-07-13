$(document).ready(function(){

	$('#id_cliente').on('change',function(){
		var id = $('#id_cliente').val();
		
		$.post('/servicio/register/'+window.idpath, {'id':id}, function(datos){
			
			alert(datos);
			$('#id_mascota').html(datos)
			
		})
		
		/*$.ajax({
			type: 'POST',
			url: '/servicio/register/'+window.idpath,
			data: {'id':id}
		})
		.done(function(lista_mascota){
			alert(a)
			$('#id_mascota').html(lista_mascota)
		})
		.fail(function(){
			alert('error a cargar las mascotas')
		})*/
	})
	
})