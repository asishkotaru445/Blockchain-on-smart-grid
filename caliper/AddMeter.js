'use strict';

const { WorkloadModuleBase } = require('@hyperledger/caliper-core');
var createdmeters=[];
class MyWorkload extends WorkloadModuleBase {
    constructor() {
        super();
    }

    async initializeWorkloadModule(workerIndex, totalWorkers, roundIndex, roundArguments, sutAdapter, sutContext) {
        await super.initializeWorkloadModule(workerIndex, totalWorkers, roundIndex, roundArguments, sutAdapter, sutContext);

        // for (let i=0; i<this.roundArguments.assets; i++) {
        //     const assetID = `${this.workerIndex}_${i}`;
        //     console.log(`Worker ${this.workerIndex}: Creating meter ${assetID}`);
        //     const request = {
        //         contractId: this.roundArguments.contractId,
        //         contractFunction: 'AddMeter',
        //         invokerIdentity: 'User1',
        //         contractArguments: [assetID,'as23df','500'],
        //         readOnly: false
        //     };

        //     await this.sutAdapter.sendRequests(request);
        // }
    }

    async submitTransaction() {
        const randomId = Math.floor(Math.random()*this.roundArguments.assets);
        const myArgs = {
            contractId: this.roundArguments.contractId,
            contractFunction: 'AddMeter',
            invokerIdentity: 'User1',
            contractArguments: [`${this.workerIndex}_${randomId}`,"createdsec_par13e12e1212e","createdh_pk23e1ar23"],
            readOnly: true
        };
        createdmeters.concat(`${this.workerIndex}_${randomId}`);
        await this.sutAdapter.sendRequests(myArgs);
    }

    async cleanupWorkloadModule() {
        for(let i=0;i<createdmeters.length();i++){
            const assetID = createdmeters[i];
            console.log(`Worker ${this.workerIndex}: Deleting meter ${assetID}`);
            const request = {
                contractId: this.roundArguments.contractId,
                contractFunction: 'RemoveMeter',
                invokerIdentity: 'User1',
                contractArguments: [assetID],
                readOnly: false
            };

            await this.sutAdapter.sendRequests(request);
        }
        // for (let i=0; i<this.roundArguments.assets; i++) {
        //     const assetID = `${this.workerIndex}_${i}`;
        //     console.log(`Worker ${this.workerIndex}: Deleting meter ${assetID}`);
        //     const request = {
        //         contractId: this.roundArguments.contractId,
        //         contractFunction: 'RemoveMeter',
        //         invokerIdentity: 'User1',
        //         contractArguments: [assetID],
        //         readOnly: false
        //     };

        //     await this.sutAdapter.sendRequests(request);
        // }
    }
}

function createWorkloadModule() {
    return new MyWorkload();
}

module.exports.createWorkloadModule = createWorkloadModule;