var crudList = angular.module('crud-owner', ['crud-services']);

crudList.controller('crudOwnerController', ['$rootScope', '$scope', '$location', 'CrudActions', 'shareDataService',
                                                             function ($rootScope, $scope, $location, CrudActions, shareDataService) {


	$scope.setOwner = function setOwner(owner){
		shareDataService.setValue('owner', owner);
		$location.path('list');
	};
}]);
