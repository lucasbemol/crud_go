var atsControllers = angular.module('crud-controllers', ['crud-services']);

/* Navigation Controller */
atsControllers.controller('NavController', ['$scope', '$rootScope', '$route', '$location', function ($scope, $rootScope, $route, $location) {
    $rootScope.route = $route;
    
	$scope.redirectToLocation = function(locationPath) {
		$location.path(locationPath);
	};

}]);