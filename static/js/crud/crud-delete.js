var crudDelete = angular.module('crud-delete', ['crud-services']);


crudDelete.controller('crudDeleteController', ['$rootScope', '$scope', '$location', '$route' ,'CrudActions', 'shareDataService',
                                                             function ($rootScope, $scope, $location, $route, CrudActions, shareDataService) {

	$scope.deleteProduct = function deleteProduct(id_product){
		//Força fechamento da Dialog
		$('#delete-modal').modal('hide');
		$('body').removeClass('modal-open');
		$('.modal-backdrop').remove();
		CrudActions.delete({ id_product : id_product},function(data){
			$scope.result = data.status;		
			if($scope.result == 'success'){
				alert('Produto excluído com sucesso!');
			}else{
				alert('Erro ao excluir produto!');
			}
			$location.path("/");
		});
	};

	$scope.deleteProductByList = function deleteProductByList(){
		//Força fechamento da Dialog
		var id_to_delete = shareDataService.getValue('id_product_to_delete');
		$('#delete-modal').modal('hide');
		$('body').removeClass('modal-open');
		$('.modal-backdrop').remove();
		CrudActions.delete({ id_product : id_to_delete},function(data){
			$scope.result = data.status;		
			if($scope.result == 'success'){
				alert('Produto excluído com sucesso!');
			}else{
				alert('Erro ao excluir produto!');
			}
			$route.reload();
		});
	};                                                              	

}]);