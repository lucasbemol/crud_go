var crudVisualize = angular.module('crud-visualize', ['crud-services']);

crudVisualize.controller('crudVisualizeController', ['$rootScope', '$scope', '$location','CrudActions', 'shareDataService',
                                                             function ($rootScope, $scope, $location,CrudActions, shareDataService) {

	CrudActions.search({ id_product : shareDataService.getValue('id_product')},function(data){
		$scope.product = data.results;
	});

	$scope.editProduct = function editProduct(){
		shareDataService.setValue('product', $scope.product);
		$location.path("edit");
	};

}]);