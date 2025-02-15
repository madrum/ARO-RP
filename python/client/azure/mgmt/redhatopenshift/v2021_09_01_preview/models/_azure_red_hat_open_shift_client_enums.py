# coding=utf-8
# --------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#   http://www.apache.org/licenses/LICENSE-2.0
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
# 
# Code generated by Microsoft (R) AutoRest Code Generator.Changes may cause incorrect behavior and will be lost if the code is regenerated.
# --------------------------------------------------------------------------

from enum import Enum, EnumMeta
from six import with_metaclass

class _CaseInsensitiveEnumMeta(EnumMeta):
    def __getitem__(self, name):
        return super().__getitem__(name.upper())

    def __getattr__(cls, name):
        """Return the enum member matching `name`
        We use __getattr__ instead of descriptors or inserting into the enum
        class' __dict__ in order to support `name` and `value` being both
        properties for enum members (which live in the class' __dict__) and
        enum members themselves.
        """
        try:
            return cls._member_map_[name.upper()]
        except KeyError:
            raise AttributeError(name)


class CreatedByType(with_metaclass(_CaseInsensitiveEnumMeta, str, Enum)):
    """The type of identity that created the resource.
    """

    USER = "User"
    APPLICATION = "Application"
    MANAGED_IDENTITY = "ManagedIdentity"
    KEY = "Key"

class EncryptionAtHost(with_metaclass(_CaseInsensitiveEnumMeta, str, Enum)):
    """EncryptionAtHost represents encryption at host state
    """

    DISABLED = "Disabled"
    ENABLED = "Enabled"

class ProvisioningState(with_metaclass(_CaseInsensitiveEnumMeta, str, Enum)):
    """ProvisioningState represents a provisioning state.
    """

    ADMIN_UPDATING = "AdminUpdating"
    CREATING = "Creating"
    DELETING = "Deleting"
    FAILED = "Failed"
    SUCCEEDED = "Succeeded"
    UPDATING = "Updating"

class SoftwareDefinedNetwork(with_metaclass(_CaseInsensitiveEnumMeta, str, Enum)):
    """SoftwareDefinedNetwork constants.
    """

    OVN_KUBERNETES = "OVNKubernetes"
    OPEN_SHIFT_SDN = "OpenShiftSDN"

class Visibility(with_metaclass(_CaseInsensitiveEnumMeta, str, Enum)):
    """Visibility represents visibility.
    """

    PRIVATE = "Private"
    PUBLIC = "Public"

class VMSize(with_metaclass(_CaseInsensitiveEnumMeta, str, Enum)):
    """VMSize represents a VM size.
    """

    STANDARD_D16_AS_V4 = "Standard_D16as_v4"
    STANDARD_D16_S_V3 = "Standard_D16s_v3"
    STANDARD_D2_S_V3 = "Standard_D2s_v3"
    STANDARD_D32_AS_V4 = "Standard_D32as_v4"
    STANDARD_D32_S_V3 = "Standard_D32s_v3"
    STANDARD_D4_AS_V4 = "Standard_D4as_v4"
    STANDARD_D4_S_V3 = "Standard_D4s_v3"
    STANDARD_D8_AS_V4 = "Standard_D8as_v4"
    STANDARD_D8_S_V3 = "Standard_D8s_v3"
    STANDARD_E16_S_V3 = "Standard_E16s_v3"
    STANDARD_E32_S_V3 = "Standard_E32s_v3"
    STANDARD_E4_S_V3 = "Standard_E4s_v3"
    STANDARD_E64_I_V3 = "Standard_E64i_v3"
    STANDARD_E64_IS_V3 = "Standard_E64is_v3"
    STANDARD_E8_S_V3 = "Standard_E8s_v3"
    STANDARD_F16_S_V2 = "Standard_F16s_v2"
    STANDARD_F32_S_V2 = "Standard_F32s_v2"
    STANDARD_F4_S_V2 = "Standard_F4s_v2"
    STANDARD_F72_S_V2 = "Standard_F72s_v2"
    STANDARD_F8_S_V2 = "Standard_F8s_v2"
    STANDARD_G5 = "Standard_G5"
    STANDARD_GS5 = "Standard_GS5"
    STANDARD_M128_MS = "Standard_M128ms"
