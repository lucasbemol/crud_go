/*
ATS custom extension of http://jsfiddle.net/NDFHg/489/
*/

var htmlEditable = angular.module('angular-html-editable', []);

htmlEditable.directive('vnEditInPlace', function () {
    return {
        restrict: 'E',
        scope: {
            value: '=',
        	editButtonCondition: '@'
        },
        template: 
        	'<span ng-bind="value"></span>'+
        	'<input ng-blur="setEdited()" ng-model="value" maxlength="40"></input>'+
        	'<br /><br />'+
        	'<button class="mar" ng-show="editing == false && {{editButtonCondition}}" ng-click="edit()" translate="vamosnessa.ats.general.button.rename"></button>'+
        	'<button class="mar" ng-show="editing == true" ng-click="setEdited()" translate="vamosnessa.ats.general.button.ok"></button>',
        link: function ($scope, element, attrs) {
            // get a reference to the input element
            var inputElement = angular.element(element.children()[1]);

            // add class to the directive
            element.addClass('edit-in-place');

            // Initially, we're not editing.
            $scope.editing = false;

            // ng-click handler to activate edit-in-place
            $scope.edit = function () {
                $scope.editing = true;

                // We control display through a class on the directive itself. See the CSS.
                element.addClass('active');

                // And we must focus the element. 
                // `angular.element()` provides a chainable array, like jQuery so to access a native DOM function, 
                // we have to reference the first element in the array.
                inputElement[0].focus();
            };
            
            $scope.setEdited = function () {
            	$scope.editing = false;
                element.removeClass('active');
            };

        }
    };
});