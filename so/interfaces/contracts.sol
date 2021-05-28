// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.7.0;

abstract contract GoodDataContract {
    //QSC
    function registerQC(bytes memory uuid, bytes memory publicKey) public virtual {
        emit QCRegistered(uuid, publicKey);
    }
    event QCRegistered(bytes uuid, bytes publicKey);

    // Using ',' to split among mpcUUIDs
    function assignQueryToMPCs(bytes memory queryUUID, bytes memory mpcUUIDs) public virtual {
        emit MPCsOfQuery(queryUUID, mpcUUIDs);
    }
    event MPCsOfQuery(bytes queryUUID, bytes mpcUUIDs);

    // Using ',' to split among mpcUUIDs
    function assignQueryToDOs(bytes memory queryUUID, bytes memory doUUIDs) public virtual {
        emit DOsOfQuery(queryUUID, doUUIDs);
    }
    event DOsOfQuery(bytes queryUUID, bytes doUUIDs);

    function submitQuery(bytes memory queryUUID, bytes memory qcUUID, int64 dataSet, uint256 fees) public virtual {
        emit QuerySubmitted(queryUUID, qcUUID, dataSet);
    }
    event QuerySubmitted(bytes queryUUID, bytes qcUUID, int64 dataSet);

    function finishQueryInOneDO(bytes memory queryUUID, bytes memory doUUID) public virtual {
        emit QueryFinishedInOneDO(queryUUID, doUUID);
    }
    event QueryFinishedInOneDO(bytes queryUUID, bytes doUUID);

    function finishQuery(bytes memory queryUUID) public virtual {
        emit QueryFinished(queryUUID);
    }
    event QueryFinished(bytes queryUUID);
    
    function submitPrediction(bytes memory queryUUID, bytes memory predictionUUID) public virtual {
        emit PredictionSubmitted(queryUUID, predictionUUID);
    }
    event PredictionSubmitted(bytes queryUUID, bytes predictionUUID);
    
    function finishPrediction(bytes memory predictionUUID) public virtual {
        emit PredictionFinished(predictionUUID);
    }
    event PredictionFinished(bytes predictionUUID);

    //DSC
    function registerDO(bytes memory uuid, bytes memory publicKey) public virtual {
        emit DORegistered(uuid, publicKey);
    }
    event DORegistered(bytes uuid, bytes publicKey);

    //MSC
    function registerMPC(bytes memory uuid, bytes memory publicKey) public virtual {
        emit MPCRegistered(uuid, publicKey);
    }
    event MPCRegistered(bytes uuid, bytes publicKey);

    //Other queryUUID
    function addLog(bytes memory uuid, bytes memory payload) public virtual {
        emit LogAdded(uuid, payload);
    }
    event LogAdded(bytes uuid, bytes payload);
}
