var crudList = angular.module('crud-list', ['crud-services']);

crudList.controller('crudListController', ['$rootScope', '$scope', '$location', 'CrudActions', 'shareDataService',
                                                             function ($rootScope, $scope, $location, CrudActions, shareDataService) {

	$scope.productsList = {};

	CrudActions.list({ owner : shareDataService.getValue('owner') }, function(data){
		$scope.productsList = data.results;
	});


	$scope.visualizeProduct = function visualizeProduct(id_product){
		shareDataService.setValue('id_product', id_product);
		$location.path('visualize');
	};


	$scope.editProduct = function editProduct(product){
		shareDataService.setValue('product', product);
		$location.path('edit');
	};

	$scope.prepareDelete = function prepareDelete(id_product_to_delete){
		shareDataService.setValue('id_product_to_delete', id_product_to_delete);
	};

	$scope.searchByName = function searchByName(){
		if($scope.searchName == "" || $scope.searchName == undefined){
			CrudActions.list(function(data){
				$scope.productsList = data.results;
			});
		}else{
			var name = '%' + $scope.searchName + '%';
			CrudActions.searchByName({ name : name, owner : shareDataService.getValue('owner') }, function(data){
				$scope.productsList = data.results;
			});
		}
	};

}]);
