var crudEdit = angular.module('crud-edit', ['crud-services']);

crudEdit.controller('crudEditController', ['$rootScope', '$scope', '$location','CrudActions', 'shareDataService',
                                                             function ($rootScope, $scope, $location,CrudActions, shareDataService) {
    
    $('.datepicker').datepicker({
 		format: 'yyyy/mm/dd',
 		language: 'pt-BR'
 	});

	$scope.product = shareDataService.getValue('product');

	$scope.editProduct = function editProduct(){
		if($scope.formEdit.$valid){
			CrudActions.update($scope.product,function(data){
				$scope.result = data.status;
				if($scope.result == 'success'){
					alert('Produto editado com sucesso!');
				}else{
					alert('Erro ao editar produto!');
				}

				$location.path('/');
			});
		}
	}                                                             	

}]);