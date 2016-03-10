/*
The MIT License (MIT)

Copyright (c) 2016 Chaabane Jalal

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

var colors = ['#F7464A', '#FDB45C', '#949FB1',
              '#4D5360', '#E67E22', '#949FB1',
              '#9B59B6', '#3498DB', '#2ECC71',
              '#D64541', '#674172', '#68C3A3',
              '#C8F7C5', '#E9D460', '#AEA8D3',
              '#BFBFBF', '#F27935', '#EF4836'];
var hightlights = {
    '#F7464A' : '#FF5A5E', '#FDB45C' : '#5AD3D1', '#949FB1' : '#A8B3C5',
    '#4D5360' : '#616774', '#E67E22' : '#D35400', '#949FB1' : '#A8B3C5',
    '#9B59B6' : '#8E44AD', '#3498DB' : '#2980B9', '#2ECC71' : '#27AE60',
    '#D64541' : '#D64520', '#674172' : '#674142', '#68C3A3' : '#68C3A1',
    '#C8F7C5' : '#C8F720', '#E9D460' : '#E9D45A', '#AEA8D3' : '#AEA8C0',
    '#BFBFBF' : '#BFBFA0', '#F27935' : '#F2790F', '#EF4836' : '#EF4830'
};

function hasStats(repository) {
    return Object.keys(repository.language_stats).length > 0
}

function loadLanguageStats(repositories) {
    var graphs = document.getElementsByClassName('lang-stats-graph');

    // Iterate over all repository canvas elements
    Array.prototype.forEach.call(graphs, function(graph, i) {
        if (hasStats(repositories[i])) {
            // Transform object of stats into array of stats
            var langStats = Object.keys(repositories[i].language_stats).map(function(key, j) {
                return {
                    value : repositories[i].language_stats[key],
                    label : key,
                    color : colors[j % colors.length],
                    hightlights : hightlights[this.color]
                };
            });
            var ctx = graph.getContext('2d');

            new Chart(ctx).Doughnut(langStats);
        }
    });
}

function loadGlobalLanguageStats(repositories) {

    // Group all project language stats into one global language stats
    var stats = repositories.reduce(function(result, repository) {
        Object.keys(repository.language_stats).forEach(function(key) {
            if (result[key] != null) {
                result[key] += repository.language_stats[key]
           } else {
               result[key] = repository.language_stats[key];
           }
           return result;
        });
        return result;
    }, {});

    // use stats to setup the graph data used by Chart.js
    var graphData = Object.keys(stats).map(function(key, i) {
        return {
            value : stats[key],
            label : key,
            color : colors[i % colors.length],
            hightlights : hightlights[this.color]
        };
    });

    var ctx = document.getElementById('lang-stats-global').getContext('2d');

    new Chart(ctx).Doughnut(graphData);
}
