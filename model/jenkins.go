package model

import (
	"fmt"
	"github.com/bndr/gojenkins"
	"time"
)

var (
	jenkins = "https://tourism.cdcyy.cn/jenkins"
	username = "admin"
	password = "admin"
)

type JenkinsManage struct {
	client *gojenkins.Jenkins
}

func NewJenkinsManage() *JenkinsManage {
	m := &JenkinsManage{}
	jks := gojenkins.CreateJenkins(nil,jenkins,username,password)
	var err error
	m.client,err = jks.Init()
	if err != nil {
		fmt.Println(err)
	}
	return m
}

func (j *JenkinsManage)GetJob(name string) (err error) {
	_,err =j.client.GetJob(name)
	return
}

func (j *JenkinsManage)CreateJobBuild(svc string,params map[string]string)(bid int,err error)  {
	_,err =j.client.GetJob(svc)
	if err!=nil {
		return
	}
	bid64,err := j.client.BuildJob(svc,params)
	return int(bid64),err
}

func (j *JenkinsManage)CreateServiceJob(name string)  (err error){
	configString := `<?xml version="1.0" encoding="utf-8"?>
<flow-definition plugin="workflow-job@2.40"> 
  <actions> 
    <org.jenkinsci.plugins.pipeline.modeldefinition.actions.DeclarativeJobAction plugin="pipeline-model-definition@1.7.2"/>  
    <org.jenkinsci.plugins.pipeline.modeldefinition.actions.DeclarativeJobPropertyTrackerAction plugin="pipeline-model-definition@1.7.2"> 
      <jobProperties/>  
      <triggers/>  
      <parameters/>  
      <options/> 
    </org.jenkinsci.plugins.pipeline.modeldefinition.actions.DeclarativeJobPropertyTrackerAction> 
  </actions>  
  <description/>  
  <keepDependencies>false</keepDependencies>  
  <properties> 
    <hudson.model.ParametersDefinitionProperty> 
      <parameterDefinitions> 
        <hudson.model.StringParameterDefinition> 
          <name>docker_repo</name>  
          <description>registry.cn-hangzhou.aliyuncs.com</description>  
          <defaultValue/>  
          <trim>false</trim> 
        </hudson.model.StringParameterDefinition>  
        <hudson.model.StringParameterDefinition> 
          <name>namespace</name>  
          <description>online-test</description>  
          <defaultValue/>  
          <trim>false</trim> 
        </hudson.model.StringParameterDefinition>  
        <hudson.model.StringParameterDefinition> 
          <name>image_tag</name>  
          <description>20201208-k222</description>  
          <defaultValue/>  
          <trim>false</trim> 
        </hudson.model.StringParameterDefinition>  
        <hudson.model.StringParameterDefinition> 
          <name>service_name</name>  
          <description>park-service</description>  
          <defaultValue/>  
          <trim>false</trim> 
        </hudson.model.StringParameterDefinition>  
        <hudson.model.StringParameterDefinition> 
          <name>gitlab</name>  
          <description>https://dev-gitlab.iuctrip.com/tesla/park-service.git</description>  
          <defaultValue/>  
          <trim>false</trim> 
        </hudson.model.StringParameterDefinition>  
        <hudson.model.StringParameterDefinition> 
          <name>tag</name>  
          <description>t1.37.1</description>  
          <defaultValue/>  
          <trim>false</trim> 
        </hudson.model.StringParameterDefinition> 
      </parameterDefinitions> 
    </hudson.model.ParametersDefinitionProperty> 
  </properties>  
  <definition class="org.jenkinsci.plugins.workflow.cps.CpsFlowDefinition" plugin="workflow-cps@2.87"> 
    <script>pipeline { 
      agent any 
      stages { 
        stage('部署') { 
          steps { 
            deleteDir() 
          } 
        } 
        stage('checkout'){ 
          steps { 
            git branch: 'develop', credentialsId: '71253ece-c038-48db-9f5e-27650109e07e', url: "${gitlab}" 
            sh ''' 
               git checkout ${tag} 
               git checkout -b ${tag##*/} 
               git show HEAD 
            ''' 
          } 
        } 
        stage('build') {
          steps { 
            sh "mvn -pl ${service_name}-spring-boot clean package" 
          } 
        } 
        stage('build images') { 
          steps { 
            sh """ 
               bash -c "${JENKINS_HOME}/scripts/deploy.sh ${WORKSPACE} ${JENKINS_HOME} ${service_name} ${docker_repo} ${namespace} ${image_tag}"
            """ 
            }
          } 
      } 
}</script>  
    <sandbox>true</sandbox> 
  </definition>  
  <triggers/>  
  <disabled>false</disabled> 
</flow-definition>`
	_,err =j.client.CreateJob(configString,name)
	if err != nil {
		fmt.Println(err)
		return
	}
	if  job,err:=j.client.GetJob(name);err!=nil{
		job.Enable()
	}
	return
}

func (j *JenkinsManage)CreateDevServiceJob(name string) (err error)  {
	configString := `<?xml version="1.0" encoding="utf-8"?>
<flow-definition plugin="workflow-job@2.40"> 
  <actions> 
    <org.jenkinsci.plugins.pipeline.modeldefinition.actions.DeclarativeJobAction plugin="pipeline-model-definition@1.7.2"/>  
    <org.jenkinsci.plugins.pipeline.modeldefinition.actions.DeclarativeJobPropertyTrackerAction plugin="pipeline-model-definition@1.7.2"> 
      <jobProperties/>  
      <triggers/>  
      <parameters/>  
      <options/> 
    </org.jenkinsci.plugins.pipeline.modeldefinition.actions.DeclarativeJobPropertyTrackerAction> 
  </actions>  
  <description/>  
  <keepDependencies>false</keepDependencies>  
  <properties> 
    <hudson.model.ParametersDefinitionProperty> 
      <parameterDefinitions> 
        <hudson.model.StringParameterDefinition> 
          <name>docker_repo</name>  
          <description>registry.cn-hangzhou.aliyuncs.com</description>  
          <defaultValue>registry.cn-hangzhou.aliyuncs.com</defaultValue>  
          <trim>false</trim> 
        </hudson.model.StringParameterDefinition>  
        <hudson.model.StringParameterDefinition> 
          <name>namespace</name>  
          <description>online-test</description>  
          <defaultValue>online-test</defaultValue>  
          <trim>false</trim> 
        </hudson.model.StringParameterDefinition>  
        <hudson.model.StringParameterDefinition> 
          <name>image_tag</name>  
          <description>20201208-k222</description>  
          <defaultValue>20201208-k222</defaultValue>  
          <trim>false</trim> 
        </hudson.model.StringParameterDefinition>  
        <hudson.model.StringParameterDefinition> 
          <name>service_name</name>  
          <description>park-service</description>  
          <defaultValue>park-service</defaultValue>  
          <trim>false</trim> 
        </hudson.model.StringParameterDefinition>  
        <hudson.model.StringParameterDefinition> 
          <name>gitlab</name>  
          <description>https://dev-gitlab.iuctrip.com/tesla/park-service.git</description>  
          <defaultValue>https://dev-gitlab.iuctrip.com/tesla/park-service.git</defaultValue>  
          <trim>false</trim> 
        </hudson.model.StringParameterDefinition>  
        <hudson.model.StringParameterDefinition> 
          <name>tag</name>  
          <description>t1.37.1</description>  
          <defaultValue>t1.37.1</defaultValue>  
          <trim>false</trim> 
        </hudson.model.StringParameterDefinition> 
      </parameterDefinitions> 
    </hudson.model.ParametersDefinitionProperty> 
  </properties>  
  <definition class="org.jenkinsci.plugins.workflow.cps.CpsFlowDefinition" plugin="workflow-cps@2.87"> 
    <script>pipeline {  
  agent any 
  stages{ 
    stage('部署') { 
      steps { 
        deleteDir()
        sh """ 
          printenv
        """
        sh 'sleep 120' 
      } 
    } 
  } 
}</script>  
    <sandbox>true</sandbox> 
  </definition>  
  <triggers/>  
  <disabled>false</disabled> 
</flow-definition>`
	_,err =j.client.CreateJob(configString,name)
	if err != nil {
		fmt.Println(err)
		return
	}
	if  job,err:=j.client.GetJob(name);err!=nil{
		job.Enable()
	}
	return
}

func (j *JenkinsManage)GetJobBuildStatus(jname string,jbid int)(running,good bool)  {
	buidl,err:=j.client.GetBuild(jname,int64(jbid))
	if err != nil {
		fmt.Println(err)
	}
	return buidl.IsRunning(),buidl.IsGood()
}


func (j *JenkinsManage)GetJobBuildId(qid int)  (bid int){
	var buildNumber int64
	for{
		task,_ := j.client.GetQueueItem(int64(qid))
		if task.Raw.Executable.Number != 0 {
			buildNumber = task.Raw.Executable.Number
			break
		}
		time.Sleep(time.Second * 1)
	}
	return int(buildNumber)
}