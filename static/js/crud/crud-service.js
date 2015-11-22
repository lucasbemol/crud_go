'use strict';

var crudServices = angular.module('crud-services', ['ngResource']);

crudServices.factory('CrudActions', ['$resource', function ($resource) {
	return $resource('/product/', {id_product : '@id_product'}, {
		'add': {
			method	: 'POST',
			url		: 'product/add/data.json',
		},
		'list': {
			method	: 'GET',
			url		: 'product/list'
		},
		'update': {
			method	: 'PUT',
			url		: 'product/:id_product/data.json'
		},
		'delete': {
			method	: 'DELETE',
			url		: 'product/delete/:id_product'
		},
		'search': {
			method	: 'GET',
			url		: 'product/search/:id_product'
		},
		'searchByName': {
			method	: 'GET',
			url		: 'product/searchByName/:name'
		}
	});
}]);


crudServices.factory('shareDataService', function() {
	var object = {};
	return {
		setValue : function(objectName, objectValue) {
			object[objectName] = objectValue;
		},
		getValue : function(objectName) {
			return object[objectName];
		},
		clearObject : function() {
			object = {};
		}
	}
});