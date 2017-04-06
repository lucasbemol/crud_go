'use strict';

var crudApp = angular.module('crudApp', [
    'crud-services',
    'ngRoute',
    'crud-list',
    'crud-visualize',
    'crud-delete',
    'crud-add',
    'crud-edit',
    'crud-owner',
    'ui.utils.masks'

]);

crudApp.config(['$routeProvider', '$locationProvider', function ($routeProvider, $locationProvider) {
    $routeProvider.
        when('/', {
        	templateUrl 		: 'owner.html',
            controller          : 'crudOwnerController'
        }).
        when('/list', {
        	templateUrl 		: 'list.html',
            controller          : 'crudListController'
        }).
        when('/home', {
        	templateUrl 		: 'list.html',
            controller          : 'crudListController'
        }).
        when('/add', {
        	templateUrl 		: 'add.html',
            controller          : 'crudAddController'

        }).
        when('/visualize', {
            templateUrl         : 'view.html',
            controller          : 'crudVisualizeController'

        }).
        when('/edit', {
        	templateUrl 		: 'edit.html',
            controller          : 'crudEditController'
        }).
        otherwise({
        	redirectTo	: '/'
        });

    $locationProvider.html5Mode(true);
}]);
