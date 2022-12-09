// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { Semver } from "@eth-optimism/contracts-bedrock/contracts/universal/Semver.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

/**
 * @title AttestationStation
 * @dev   Contract for creating attestations.
 * @notice This contract is not yet audited.
 */
contract AttestationStation is Semver {
    /**
     * @notice Struct representing data that is being attested
     *
     * @custom:field about                    Address being attested about (not creator/msg.sender)
     * @custom:field key                      A bytes32 key for the attestation.
     * @custom:field val                      The attestation as arbitrary bytes
     */
    struct AttestationData {
        address about;
        bytes32 key;
        bytes val;
    }

    /**
     * @notice Maps addresses to attestations
     * @dev addresses map to attestations map of
     *      `about` addresses to key/values
     *      key/values are a map of bytes32 to bytes
     */
    mapping(address => mapping(address => mapping(bytes32 => bytes))) public attestations;

    /**
     * @notice Emitted when Attestation is created
     *
     * @param creator Address that attested.
     * @param about  Address attestation is about.
     * @param key Key of the attestation.
     * @param val Value of the attestation.
     */
    event AttestationCreated(
        address indexed creator,
        address indexed about,
        bytes32 indexed key,
        bytes val
    );

    constructor() Semver(0, 0, 1) {}

    /**
     * @notice  Attest to the given data.
     * @dev     Attests to the given data from the sender.
     * @param   _about  The address of the attestation subject.
     * @param   _key  The key of the attestation.
     * @param   _val  The value of the attestation.
     */
    function _attest(
        address _about,
        bytes32 _key,
        bytes memory _val
    ) private {
        attestations[msg.sender][_about][_key] = _val;
        emit AttestationCreated(msg.sender, _about, _key, _val);
    }

    /**
     * @notice  Attest to the given data.
     * @dev     Attests to the given data from the sender.
     * @param   _attestations  The array of attestation data.
     */
    function attest(AttestationData[] memory _attestations) public {
        uint256 length = _attestations.length;
        for (uint256 i = 0; i < length; ) {
            _attest(_attestations[i].about, _attestations[i].key, _attestations[i].val);
            unchecked {
                ++i;
            }
        }
    }
}
