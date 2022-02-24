import { distinct, maxIndex, stringComparator } from '/tbc/core/utils.js';
import { actionColors } from './color_settings.js';
import { ResultComponent } from './result_component.js';
const dpsColor = '#ed5653';
const manaColor = '#2E93fA';
export class Timeline extends ResultComponent {
    constructor(config) {
        config.rootCssClass = 'timeline-root';
        super(config);
        this.resultData = null;
        this.rendered = false;
        this.rootElem.innerHTML = `
		<div class="timeline-disclaimer">
			<span class="timeline-warning fa fa-exclamation-triangle"></span>
			<span class="timeline-warning-description">Timeline data visualizes only 1 sim iteration.</span>
			<div class="timeline-run-again-button sim-button">SIM 1 ITERATION</span>
		</div>
		<div class="timeline-plots-container">
			<div class="timeline-plot dps-resources-plot"></div>
			<div class="timeline-plot cooldowns-plot"></div>
		</div>
		`;
        const runAgainButton = this.rootElem.getElementsByClassName('timeline-run-again-button')[0];
        runAgainButton.addEventListener('click', event => {
            (window.opener || window.parent).postMessage('runOnce', '*');
        });
        this.dpsResourcesPlotElem = this.rootElem.getElementsByClassName('dps-resources-plot')[0];
        this.dpsResourcesPlot = new ApexCharts(this.dpsResourcesPlotElem, {
            chart: {
                type: 'line',
                foreColor: 'white',
                id: 'dpsResources',
                animations: {
                    enabled: false,
                },
                height: '100%',
            },
            colors: [
                dpsColor,
                manaColor,
            ],
            series: [],
            xaxis: {
                title: {
                    text: 'Time (s)',
                },
                type: 'datetime',
            },
            yaxis: {},
            noData: {
                text: 'Waiting for data...',
            },
            stroke: {
                width: 2,
                curve: 'straight',
            },
        });
        this.cooldownsPlotElem = this.rootElem.getElementsByClassName('cooldowns-plot')[0];
        //this.cooldownsPlot = new ApexCharts(this.cooldownsPlotElem, {
        //	chart: {
        //		type: 'rangeBar',
        //		foreColor: 'white',
        //		id: 'cooldowns',
        //		animations: {
        //			enabled: false,
        //		},
        //		height: '50%',
        //	},
        //	series: [], // Set dynamically
        //	noData: {
        //		text: 'Waiting for data...',
        //	},
        //});
    }
    onSimResult(resultData) {
        this.resultData = resultData;
        if (this.rendered) {
            this.updatePlot();
        }
    }
    updatePlot() {
        const players = this.resultData.result.getPlayers(this.resultData.filter);
        if (players.length != 1) {
            return;
        }
        const player = players[0];
        const duration = this.resultData.result.result.firstIterationDuration || 1;
        let manaLogs = player.manaChangedLogs;
        let dpsLogs = player.dpsLogs;
        let mcdLogs = player.majorCooldownLogs;
        let mcdAuraLogs = player.majorCooldownAuraUptimeLogs;
        if (dpsLogs.length == 0) {
            return;
        }
        const maxDps = dpsLogs[maxIndex(dpsLogs.map(l => l.dps))].dps;
        const dpsAxisMax = (Math.floor(maxDps / 100) + 1) * 100;
        // Figure out how much to vertically offset cooldown icons, for cooldowns
        // used very close to each other. This is so the icons don't overlap.
        const MAX_ALLOWED_DIST = 10;
        const cooldownIconOffsets = mcdLogs.map((mcdLog, mcdIdx) => mcdLogs.filter((cdLog, cdIdx) => (cdIdx < mcdIdx) && (cdLog.timestamp > mcdLog.timestamp - MAX_ALLOWED_DIST)).length);
        const distinctMcdAuras = distinct(mcdAuraLogs, (a, b) => a.aura.equalsIgnoringTag(b.aura));
        // Sort by name so auras keep their same colors even if timings change.
        distinctMcdAuras.sort((a, b) => stringComparator(a.aura.name, b.aura.name));
        const mcdAuraColors = mcdAuraLogs.map(mcdAuraLog => actionColors[distinctMcdAuras.findIndex(dAura => dAura.aura.equalsIgnoringTag(mcdAuraLog.aura))]);
        const showMana = manaLogs.length > 0;
        const maxMana = showMana ? manaLogs[0].valueBefore : 0;
        let options = {
            series: [{
                    name: 'DPS',
                    type: 'line',
                    data: dpsLogs.map(log => {
                        return {
                            x: this.toDatetime(log.timestamp),
                            y: log.dps,
                        };
                    }),
                }],
            xaxis: {
                min: this.toDatetime(0).getTime(),
                max: this.toDatetime(duration).getTime(),
                type: 'datetime',
                tickAmount: 10,
                decimalsInFloat: 1,
                labels: {
                    show: true,
                    formatter: (defaultValue, timestamp) => {
                        return (timestamp / 1000).toFixed(1);
                    },
                },
                title: {
                    text: 'Time (s)',
                },
            },
            yaxis: [
                {
                    color: dpsColor,
                    seriesName: 'DPS',
                    min: 0,
                    max: dpsAxisMax,
                    tickAmount: 10,
                    decimalsInFloat: 0,
                    title: {
                        text: 'DPS',
                        style: {
                            color: dpsColor,
                        },
                    },
                    axisBorder: {
                        show: true,
                        color: dpsColor,
                    },
                    axisTicks: {
                        color: dpsColor,
                    },
                    labels: {
                        minWidth: 30,
                        style: {
                            colors: [dpsColor],
                        },
                    },
                },
            ],
            annotations: {
                position: 'back',
                xaxis: mcdAuraLogs.map((log, i) => {
                    return {
                        x: this.toDatetime(log.gainedAt).getTime(),
                        x2: this.toDatetime(log.fadedAt).getTime(),
                        fillColor: mcdAuraColors[i],
                    };
                }),
                points: mcdLogs.map((log, i) => {
                    return {
                        x: this.toDatetime(log.timestamp).getTime(),
                        y: 0,
                        image: {
                            path: log.cooldownId.iconUrl,
                            width: 20,
                            height: 20,
                            offsetY: cooldownIconOffsets[i] * -25,
                        },
                    };
                }),
            },
            tooltip: {
                enabled: true,
                custom: (data) => {
                    if (data.seriesIndex == 0) {
                        // DPS
                        const log = dpsLogs[data.dataPointIndex];
                        return `<div class="timeline-tooltip dps">
							<div class="timeline-tooltip-header">
								<span class="bold">${log.timestamp.toFixed(2)}s</span>
							</div>
							<div class="timeline-tooltip-body">
								<ul class="timeline-dps-events">
									${log.damageLogs.map(damageLog => {
                            let iconElem = '';
                            if (damageLog.cause.iconUrl) {
                                iconElem = `<img class="timeline-tooltip-icon" src="${damageLog.cause.iconUrl}">`;
                            }
                            return `
										<li>
											${iconElem}
											<span>${damageLog.cause.name}:</span>
											<span class="series-color">${damageLog.resultString()}</span>
										</li>`;
                        }).join('')}
								</ul>
								<div class="timeline-tooltip-body-row">
									<span class="series-color">DPS: ${log.dps.toFixed(2)}</span>
								</div>
							</div>
							${log.activeAuras.length == 0 ? '' : `
								<div class="timeline-tooltip-auras">
									<div class="timeline-tooltip-body-row">
										<span class="bold">Active Auras</span>
									</div>
									<ul class="timeline-active-auras">
										${log.activeAuras.map(auraLog => {
                            let iconElem = '';
                            if (auraLog.aura.iconUrl) {
                                iconElem = `<img class="timeline-tooltip-icon" src="${auraLog.aura.iconUrl}">`;
                            }
                            return `
											<li>
												${iconElem}
												<span>${auraLog.aura.name}</span>
											</li>`;
                        }).join('')}
									</ul>
								</div>`}
						</div>`;
                    }
                    else if (data.seriesIndex == 1) {
                        // Mana
                        const log = manaLogs[data.dataPointIndex];
                        return `<div class="timeline-tooltip mana">
							<div class="timeline-tooltip-header">
								<span class="bold">${log.timestamp.toFixed(2)}s</span>
							</div>
							<div class="timeline-tooltip-body">
								<div class="timeline-tooltip-body-row">
									<span class="series-color">Before: ${log.valueBefore.toFixed(1)} (${(log.valueBefore / maxMana * 100).toFixed(0)}%)</span>
								</div>
								<ul class="timeline-mana-events">
									${log.logs.map(manaChangedLog => {
                            let iconElem = '';
                            if (manaChangedLog.cause.iconUrl) {
                                iconElem = `<img class="timeline-tooltip-icon" src="${manaChangedLog.cause.iconUrl}">`;
                            }
                            return `
										<li>
											${iconElem}
											<span>${manaChangedLog.cause.name}:</span>
											<span class="series-color">${manaChangedLog.resultString()}</span>
										</li>`;
                        }).join('')}
								</ul>
								<div class="timeline-tooltip-body-row">
									<span class="series-color">After: ${log.valueAfter.toFixed(1)} (${(log.valueAfter / maxMana * 100).toFixed(0)}%)</span>
								</div>
							</div>
							${log.activeAuras.length == 0 ? '' : `
								<div class="timeline-tooltip-auras">
									<div class="timeline-tooltip-body-row">
										<span class="bold">Active Auras</span>
									</div>
									<ul class="timeline-active-auras">
										${log.activeAuras.map(auraLog => {
                            let iconElem = '';
                            if (auraLog.aura.iconUrl) {
                                iconElem = `<img class="timeline-tooltip-icon" src="${auraLog.aura.iconUrl}">`;
                            }
                            return `
											<li>
												${iconElem}
												<span>${auraLog.aura.name}</span>
											</li>`;
                        }).join('')}
									</ul>
								</div>`}
						</div>`;
                    }
                }
            },
            chart: {
                events: {
                    beforeResetZoom: () => {
                        return {
                            xaxis: {
                                min: this.toDatetime(0),
                                max: this.toDatetime(duration),
                            },
                        };
                    },
                },
            },
        };
        if (showMana) {
            options.series.push({
                name: 'Mana',
                type: 'line',
                data: manaLogs.map(log => {
                    return {
                        x: this.toDatetime(log.timestamp),
                        y: log.valueAfter,
                    };
                }),
            });
            options.yaxis.push({
                seriesName: 'Mana',
                opposite: true,
                min: 0,
                max: maxMana,
                tickAmount: 10,
                title: {
                    text: 'Mana',
                    style: {
                        color: manaColor,
                    },
                },
                axisBorder: {
                    show: true,
                    color: manaColor,
                },
                axisTicks: {
                    color: manaColor,
                },
                labels: {
                    minWidth: 30,
                    style: {
                        colors: [manaColor],
                    },
                    formatter: (val) => {
                        const v = parseFloat(val);
                        return `${v.toFixed(0)} (${(v / maxMana * 100).toFixed(0)}%)`;
                    },
                },
            });
        }
        this.dpsResourcesPlot.updateOptions(options);
        //this.cooldownsPlot.updateOptions({
        //	series: [
        //		{
        //			name: 'Lightning Bolt',
        //			data: [
        //				{
        //					x: 'GCD',
        //					y: [0, 40],
        //				},
        //				{
        //					x: 'GCD',
        //					y: [60, 100],
        //				},
        //			],
        //		},
        //		{
        //			name: 'Chain Lightning',
        //			data: [
        //				{
        //					x: 'GCD',
        //					y: [0, 40],
        //				},
        //				{
        //					x: 'GCD',
        //					y: [60, 100],
        //				},
        //			],
        //		},
        //		{
        //			name: 'Bloodlust',
        //			data: [
        //				{
        //					x: 'Cooldowns',
        //					y: [0, 40],
        //				},
        //				{
        //					x: 'Cooldowns',
        //					y: [60, 100],
        //				},
        //			],
        //		},
        //		{
        //			name: 'Innervate',
        //			data: [
        //				{
        //					x: 'Cooldowns',
        //					y: [30, 70],
        //				},
        //				{
        //					x: 'Cooldowns',
        //					y: [150, 200],
        //				},
        //			],
        //		},
        //	],
        //	xaxis: {
        //		min: this.toDatetime(0),
        //		max: this.toDatetime(duration),
        //		tickAmount: 10,
        //		decimalsInFloat: 1,
        //		labels: {
        //			show: true,
        //		},
        //	},
        //	yaxis: {
        //		title: {
        //			text: 'Cooldowns',
        //		},
        //		labels: {
        //			minWidth: 30,
        //		},
        //	},
        //	plotOptions: {
        //		bar: {
        //			horizontal: true,
        //			barHeight: '80%',
        //		},
        //	},
        //	stroke: {
        //		width: 1,
        //	},
        //	fill: {
        //		type: 'solid',
        //		opacity: 0.6,
        //	},
        //	tooltip: {
        //		enabled: true,
        //	},
        //	chart: {
        //		events: {
        //			beforeResetZoom: () => {
        //				return {
        //					xaxis: {
        //						min: this.toDatetime(0),
        //						max: this.toDatetime(duration),
        //					},
        //				};
        //			},
        //		},
        //	},
        //});
    }
    render() {
        setTimeout(() => {
            this.dpsResourcesPlot.render();
            //this.cooldownsPlot.render();
            this.rendered = true;
            if (this.resultData != null) {
                this.updatePlot();
            }
        }, 300);
    }
    toDatetime(timestamp) {
        return new Date(timestamp * 1000);
    }
}
