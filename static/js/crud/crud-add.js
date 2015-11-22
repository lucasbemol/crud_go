var crudAdd = angular.module('crud-add', ['crud-services']);

crudAdd.controller('crudAddController', ['$rootScope', '$scope', '$location','CrudActions', 'shareDataService',
                                                             function ($rootScope, $scope, $location,CrudActions, shareDataService) {
 	$('.datepicker').datepicker({
 		format: 'yyyy/mm/dd',
 		language: 'pt-BR'
 	});
 
	$scope.addProduct = function addProduct(){
		if($scope.formAdd.$valid){
			CrudActions.add($scope.product,function(data){
				$scope.result = data.status;
				if($scope.result == 'success'){
					alert('Produto adicionado com sucesso!');
				}else{
					alert('Erro ao adicionar produto: ' + $scope.result);
				}
				$location.path("/");
			});
		}
	}                                                             	

}]);